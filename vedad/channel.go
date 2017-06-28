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
	}
	return mc, err
}

func (c *Channel) StartChannel() {
	var mc collectors.Collectd
	mc, err := c.getconn()
	if err != nil {
		c.Close()
		return
	}
	ticker := time.NewTicker(time.Duration(c.meta.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			data, err := mc.Start()
			c.ctx.vedad.logf(LOG_INFO, "work(%s,%s): start to collect data", c.meta.Address, c.name)
			if err != nil {
				c.ctx.vedad.logf(LOG_ERROR, "work(%s,%s): fail to collect data : %s", c.meta.Address, c.name, err.Error())
			}
			c.ctx.vedad.pushinfluxChan <- &data
		case <-c.exitChan:
			goto exit
		}
	}
exit:
	ticker.Stop()
}

func (c *Channel) Close() error {
	close(c.exitChan)
	return nil
}
