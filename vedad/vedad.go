package vedad

import (
	"errors"
	"log"
	"os"
	"sync"
	"sync/atomic"

	//_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	goutil "github.com/hawkingrei/golang_util"
	"github.com/hawkingrei/veda/collectors"
	"github.com/hawkingrei/veda/internal/lg"
	client "github.com/influxdata/influxdb/client/v2"
	"github.com/mkevac/debugcharts"
)

type VEDAD struct {
	sync.RWMutex
	opts           atomic.Value
	waitGroup      goutil.WaitGroupWrapper
	exitChan       chan int
	pushinfluxChan chan *collectors.CollectData
	topicMap       map[string]*Topic
	c              client.Client
}

func New(opts *Options) (v *VEDAD, err error) {
	if opts.Logger == nil {
		opts.Logger = log.New(os.Stderr, opts.LogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	username := ""
	password := ""
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.1.1.89:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		return v, err
	}
	v = &VEDAD{
		c:              c,
		topicMap:       make(map[string]*Topic),
		exitChan:       make(chan int),
		pushinfluxChan: make(chan *collectors.CollectData, 100000000),
	}
	opts.logLevel, err = lg.ParseLogLevel(opts.LogLevel, opts.Verbose)
	if err != nil {
		v.logf(LOG_FATAL, "%s", err)
		os.Exit(1)
	}
	v.swapOpts(opts)
	return
}

func (v *VEDAD) getOpts() *Options {
	return v.opts.Load().(*Options)
}

func (v *VEDAD) swapOpts(opts *Options) {
	v.opts.Store(opts)
}
func (v *VEDAD) Loadmeta(meta Meta) error {
	for topicName, services := range meta.Topics {
		topic := v.GetTopic(topicName)
		for servername, service := range services {
			topic.GetChannel(servername, service)
		}
	}
	return nil
}
func (v *VEDAD) Main() {
	v.waitGroup.Wrap(func() { v.ToInfluxdb() })
	router := gin.Default()
	debugcharts.GinDebugRouter(router)
	router.Run(":8080")
	//v.waitGroup.Wrap(func() { http.ListenAndServe("localhost:6060", nil) })
	//v.waitGroup.Wrap(func() { v.lookupLoop() })
	//if v.getOpts().StatsdAddress != "" {
	//	v.waitGroup.Wrap(func() { v.statsdLoop() })
	//}
}

func (v *VEDAD) Exit() {
	v.logf(LOG_INFO, "VEDAD: closing topics")
	for _, vv := range v.topicMap {
		vv.Exit()
	}
	close(v.exitChan)
	v.waitGroup.Wait()
}

// GetTopic performs a thread safe operation
// to return a pointer to a Topic object (potentially new)
func (v *VEDAD) GetTopic(topicName string) *Topic {
	// most likely, we already have this topic, so try read lock first.
	v.RLock()
	t, ok := v.topicMap[topicName]
	v.RUnlock()
	if ok {
		return t
	}

	v.Lock()
	t, ok = v.topicMap[topicName]
	if ok {
		v.Unlock()
		return t
	}
	t = NewTopic(topicName, &context{v})
	v.topicMap[topicName] = t
	v.logf(LOG_INFO, "TOPIC(%s): created", t.name)
	v.Unlock()
	return t
}

// GetExistingTopic gets a topic only if it exists
func (v *VEDAD) GetExistingTopic(topicName string) (*Topic, error) {
	v.RLock()
	defer v.RUnlock()
	topic, ok := v.topicMap[topicName]
	if !ok {
		return nil, errors.New("topic does not exist")
	}
	return topic, nil
}

// DeleteExistingTopic removes a topic only if it exists
func (v *VEDAD) DeleteExistingTopic(topicName string) error {
	v.RLock()
	topic, ok := v.topicMap[topicName]
	if !ok {
		v.RUnlock()
		return errors.New("topic does not exist")
	}
	v.RUnlock()

	// delete empties all channels and the topic itself before closing
	// (so that we dont leave any messages around)
	//
	// we do this before removing the topic from map below (with no lock)
	// so that any incoming writes will error and not create a new topic
	// to enforce ordering
	topic.Exit()

	v.Lock()
	delete(v.topicMap, topicName)
	v.Unlock()

	return nil
}

func (v *VEDAD) ToInfluxdb() {
	for {
		select {
		case c := <-v.pushinfluxChan:
			err := v.putdata(*c)
			if err != nil {
				//v.pushinfluxChan <- c
				v.logf(LOG_ERROR, "VEDAD:(%s %s) fail to write data into influxdb : %s", c.Name, c.Tags, err.Error())
			} else {
				//v.logf(LOG_DEBUG, "VEDAD: succeed to write data into influxdb")
			}

		case <-v.exitChan:
			goto exit
		}
	}
exit:
	v.logf(LOG_DEBUG, "VEDAD: ToInfluxdb exit")
}

func (v *VEDAD) putdata(data collectors.CollectData) error {
	MyDB := "square_holes"
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "us",
	})
	if err != nil {
		return err
	}

	// Create a point and add to batch
	pt, err := client.NewPoint(data.Name, data.Tags, data.Data, data.T)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := v.c.Write(bp); err != nil {
		return err
	}
	return nil
}
