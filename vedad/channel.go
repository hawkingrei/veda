package vedad

import (
	"sync"
	"time"

	goutil "github.com/hawkingrei/golang_util"
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
	waitGroup goutil.WaitGroupWrapper

	exitFlag  int32
	exitChan  chan int
	exitMutex sync.RWMutex
}

func NewChannel(topicName string, channelName string, channelsMeta ChannelsMeta, ctx *context) *Channel {
	c := &Channel{
		topicName: topicName,
		name:      channelName,
		meta:      channelsMeta,
		ctx:       ctx,
	}
	c.waitGroup.Wrap(func() { c.StartChannel() })
	return c
}

func (c *Channel) getconn() collectors.Collectd {
	var mc collectors.Collectd
	switch c.topicName {
	case "memcache":
		mc, _ = collectors.GetMemcacheConn(c.meta.Address, c.meta.Name)
	}
	return mc
}

func (c *Channel) StartChannel() {
	var mc collectors.Collectd
	ticker := time.NewTicker(time.Duration(c.meta.Interval) * time.Second)
	mc = c.getconn()
	for {
		select {
		case <-ticker.C:
			data, err := mc.Start()
			if err == nil {
				c.ctx.vedad.pushinfluxChan <- &data
			}

		case <-c.exitChan:
			return
		}
	}
}

func (c *Channel) Close() error {
	c.waitGroup.Wait()
	return nil
}
