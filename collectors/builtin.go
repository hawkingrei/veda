package collectors

import (
	"time"
)

type CollectData struct {
	Name string
	Tags map[string]string
	Data map[string]interface{}
	T    time.Time
}

type Collectd interface {
	Start() (CollectData, error)
	convertCollectData(string) CollectData
}
