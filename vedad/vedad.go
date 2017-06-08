package vedad

import (
	goutil "github.com/hawkingrei/golang_util"
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
	return &VEDAD{
		topicMap: make(map[string]*Topic),
		exitChan: make(chan int),
	}
}

func (n *VEDAD) getOpts() *Options {
	return n.opts.Load().(*Options)
}

func (n *VEDAD) swapOpts(opts *Options) {
	n.opts.Store(opts)
}

func (n *VEDAD) Exit() {

}
