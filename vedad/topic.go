package vedad

import (
	"sync"
	"sync/atomic"
)

type Topic struct {
	messageCount uint64

	sync.RWMutex
	name     string
	exitFlag int32

	deleteCallback func(*Topic)
	deleter        sync.Once
}

func (t *Topic) Exiting() bool {
	return atomic.LoadInt32(&t.exitFlag) == 1
}
