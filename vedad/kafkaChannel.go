package vedad

import "sync"
import "github.com/hawkingrei/veda/collectors"

type KafkaChannel struct {
	topicName string
	name      string
	ctx       *context
	meta      ChannelsMeta

	exitFlag  int32
	exitChan  chan int
	exitMutex sync.RWMutex

	kafkaApp *collectors.AppConnection
}

func NewKafkaChannel(topicName string, channelName string, channelsMeta ChannelsMeta, ctx *context) *KafkaChannel {
	c := &KafkaChannel{
		exitChan:  make(chan int),
		topicName: topicName,
		name:      channelName,
		meta:      channelsMeta,
		ctx:       ctx,
	}
	c.ctx.vedad.waitGroup.Wrap(func() { c.StartChannel() })
	return c
}

func (c *KafkaChannel) StartChannel() {
	var err error
	c.kafkaApp, err = collectors.GetAppConn(c.meta.Addresses, &c.ctx.vedad.pushinfluxChan)
	if err != nil {
		c.ctx.vedad.logf(LOG_ERROR, "work(%s): fail to get connect : %s", c.name, err.Error())
		c.Close()
		return
	}
	c.ctx.vedad.waitGroup.Wrap(func() { c.kafkaApp.Start() })
}

func (c *KafkaChannel) Close() error {
	c.ctx.vedad.logf(LOG_INFO, "CHANNEL(%s): closing", c.name)
	c.kafkaApp.Close()
	return nil
}
