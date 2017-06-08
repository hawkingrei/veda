package vedad

import (
	"sync"
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

	exitFlag  int32
	exitMutex sync.RWMutex
}
