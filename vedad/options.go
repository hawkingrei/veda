package vedad

import (
	"github.com/hawkingrei/veda/internal/lg"
	//"log"
)

type Meta struct {
	Topics map[string]Channels
}

type Channels map[string]ChannelsMeta

type ChannelsMeta struct {
	Name      string
	Address   string
	Addresses []string
	Interval  int
	Username  string
	Password  string
	Db        string
}

type Options struct {
	LogLevel  string `flag:"log-level"`
	LogPrefix string `flag:"log-prefix"`
	Verbose   bool   `flag:"verbose"`
	Logger    Logger
	logLevel  lg.LogLevel

	Timeout      int
	WorkSize     int
	SendWorkSize int
}

func NewOptions() *Options {
	return &Options{
		LogLevel: "debug",
		Timeout:  1,
	}
}
