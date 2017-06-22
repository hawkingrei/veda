package vedad

import (
	goutil "github.com/hawkingrei/golang_util"
	"github.com/hawkingrei/veda/collectors"
	"sync"
	"time"
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

func (c *Channel) StartChannel() {
	var mc collectors.Collectd
	ticker := time.NewTicker(3 * time.Second)
	mc, _ = collectors.GetMemcacheConn("10.1.1.96:11211")
	for {
		select {
		case <-ticker.C:
			mc.Start()
		}
	}
}

func (c *Channel) Close() error {
	c.waitGroup.Wait()
	return nil
}
