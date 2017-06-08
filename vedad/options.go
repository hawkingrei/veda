package vedad

import (
	"github.com/hawkingrei/veda/internal/lg"
	//"log"
)

type Config struct {
	Redis map[string]Server
	//Memcache map[string]Server
}

type Server struct {
	Address string
}

type Options struct {
	LogLevel string `flag:"log-level"`
	Verbose  bool   `flag:"verbose"`
	Logger   Logger
	logLevel lg.LogLevel
	config   Config

	WorkSize     int
	SendWorkSize int
}

func NewOptions() *Options {
	return &Options{
		LogLevel: "debug",
	}
}
