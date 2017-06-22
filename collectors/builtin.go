package collectors

type CollectData struct {
	Label string
	Data  map[string]float64
}

type Collectd interface {
	Start() CollectData
}
