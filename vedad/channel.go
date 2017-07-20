package vedad

import (
	"sync"
	"time"

	"github.com/hawkingrei/veda/collectors"
)

type Consumer interface {
	UnPause()
	Pause()
	Close() error
	TimedOutMessage()
	Empty()
}

type Channel struct {
	topicName string
	name      string
	ctx       *context
	meta      ChannelsMeta

	exitFlag  int32
	exitChan  chan int
	exitMutex sync.RWMutex
}

func NewChannel(topicName string, channelName string, channelsMeta ChannelsMeta, ctx *context) *Channel {
	c := &Channel{
		exitChan:  make(chan int),
		topicName: topicName,
		name:      channelName,
		meta:      channelsMeta,
		ctx:       ctx,
	}
	c.ctx.vedad.waitGroup.Wrap(func() { c.StartChannel() })
	return c
}

func (c *Channel) getconn() (mc collectors.Collectd, err error) {
	switch c.topicName {
	case "memcache":
		c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start", "memcache", c.name)
		mc, err = collectors.GetMemcacheConn(c.meta.Address, c.name)
	case "redis":
		c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start", "redis", c.name)
		mc, err = collectors.GetRedisConn(c.meta.Address, c.name)
	case "zookeeper":
		c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start", "zookeeper", c.name)
		mc, err = collectors.GetZookeeperConn(c.meta.Address, c.name)
	case "beansdb":
		c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start", "beansdb", c.name)
		mc, err = collectors.GetBeansdbConn(c.meta.Address, c.name)
	case "nginx":
		c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start", "nginx", c.name)
		mc, err = collectors.GetNginxConn(c.meta.Address, c.name)
	case "twemproxy":
		c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start", "twemproxy", c.name)
		mc, err = collectors.GetTwemproxyConn(c.meta.Address, c.name)
	case "mysql":
		c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start", "mysql", c.name)
		mc, err = collectors.GetMysqlConn(c.meta.Address, c.meta.Username, c.meta.Password, c.meta.Db, c.name)
	}
	return mc, err
}

func (c *Channel) StartChannel() {
	var mc collectors.Collectd
	mc, err := c.getconn()
	if err != nil {
		c.ctx.vedad.logf(LOG_ERROR, "work(%s): fail to get connect : %s", c.name, err.Error())
		c.Close()
		return
	}
	//c.ctx.vedad.logf(LOG_DEBUG, "work(%s,%s): create to collect data", c.meta.Address, c.name)
	ticker := time.NewTicker(time.Duration(c.meta.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			data, err := mc.Start()
			//c.ctx.vedad.logf(LOG_DEBUG, "work(%s,%s): start to collect data", c.meta.Address, c.name)
			if err != nil {
				c.ctx.vedad.logf(LOG_ERROR, "work(%s,%s): fail to collect data : %s", c.meta.Address, c.name, err.Error())
				continue
			}
			c.ctx.vedad.pushinfluxChan <- &data
			continue
		case <-c.exitChan:
			goto exit
		}
	}
exit:
	ticker.Stop()
}

func (c *Channel) Close() error {
	c.ctx.vedad.logf(LOG_INFO, "CHANNEL(%s): closing", c.name)
	close(c.exitChan)
	return nil
}
