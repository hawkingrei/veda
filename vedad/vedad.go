package vedad

import (
	"errors"
	goutil "github.com/hawkingrei/golang_util"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

type VEDAD struct {
	sync.RWMutex
	opts      atomic.Value
	waitGroup goutil.WaitGroupWrapper
	exitChan  chan int
	topicMap  map[string]*Topic
}

func New(opts *Options) *VEDAD {
	if opts.Logger == nil {
		opts.Logger = log.New(os.Stderr, opts.LogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	v := &VEDAD{
		topicMap: make(map[string]*Topic),
		exitChan: make(chan int),
	}
	v.swapOpts(opts)
	return v
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

	//ctx := &context{v}

	//v.waitGroup.Wrap(func() { v.queueScanLoop() })
	//v.waitGroup.Wrap(func() { v.lookupLoop() })
	//if v.getOpts().StatsdAddress != "" {
	//	v.waitGroup.Wrap(func() { v.statsdLoop() })
	//}
}

func (v *VEDAD) Exit() {
	v.logf(LOG_INFO, "VEDAD: closing topics")
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
