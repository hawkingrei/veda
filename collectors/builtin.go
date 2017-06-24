package collectors

import (
	"time"
)

type CollectData struct {
	Name string
	Tags map[string]string
	Data map[string]float64
	T    time.Time
}

type Collectd interface {
	Start() CollectData
}
