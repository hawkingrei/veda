package collector

type Collectd interface {
	Get() error
}
