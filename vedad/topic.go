package vedad

import (
	"errors"
	"sync"
	"sync/atomic"

	goutil "github.com/hawkingrei/golang_util"
)

type Topic struct {
	messageCount uint64

	sync.RWMutex
	name     string
	exitFlag int32

	deleter   sync.Once
	waitGroup goutil.WaitGroupWrapper

	exitChan chan int

	channelMap map[string]*Channel
	taskMap    map[string]ChannelsMeta
	ctx        *context
}

// Topic constructor
func NewTopic(topicName string, ctx *context) *Topic {
	t := &Topic{
		name:       topicName,
		channelMap: make(map[string]*Channel),
		exitChan:   make(chan int),
		ctx:        ctx,
	}

	return t
}

func (t *Topic) Start() {
	//for {
	//	switch {
	//
	//	}
	//}
}

func (t *Topic) Exiting() bool {
	return atomic.LoadInt32(&t.exitFlag) == 1
}

// this expects the caller to handle locking
func (t *Topic) GetChannel(channelName string, channelsMeta ChannelsMeta) *Channel {
	channel, ok := t.channelMap[channelName]
	if !ok {
		channel = NewChannel(t.name, channelName, channelsMeta, t.ctx)
		t.channelMap[channelName] = channel
		t.ctx.vedad.logf(LOG_INFO, "TOPIC(%s): new channel(%s)", t.name, channel.name)
		//t.ctx.vedad.waitGroup(func() {})
		return channel
	}
	return channel
}

func (t *Topic) Exit() error {
	if !atomic.CompareAndSwapInt32(&t.exitFlag, 0, 1) {
		return errors.New("exiting")
	}

	t.ctx.vedad.logf(LOG_INFO, "TOPIC(%s): closing", t.name)

	close(t.exitChan)

	// synchronize the close of messagePump()
	t.waitGroup.Wait()

	// close all the channels
	for _, channel := range t.channelMap {
		err := channel.Close()
		if err != nil {
			// we need to continue regardless of error to close all the channels
			t.ctx.vedad.logf(LOG_ERROR, "channel(%s) close - %s", channel.name, err)
		}
	}

	// write anything leftover to disk
	return nil
}

func (t *Topic) messagePump() {

}
