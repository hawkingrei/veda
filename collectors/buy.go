package collectors

import (
	"time"
)

type Buy struct {
	Timestamp int64 `json:"timestamp"`
	Instance  struct {
		App  string `json:"app"`
		Pid  int    `json:"pid"`
		Host string `json:"host"`
	} `json:"instance"`
	Metrics struct {
		Version string `json:"version"`
		Gauges  struct {
			JvmBuffersDirectCapacity struct {
				Value int `json:"value"`
			} `json:"jvm.buffers.direct.capacity"`
			JvmBuffersDirectCount struct {
				Value int `json:"value"`
			} `json:"jvm.buffers.direct.count"`
			JvmBuffersDirectUsed struct {
				Value int `json:"value"`
			} `json:"jvm.buffers.direct.used"`
			JvmBuffersMappedCapacity struct {
				Value int `json:"value"`
			} `json:"jvm.buffers.mapped.capacity"`
			JvmBuffersMappedCount struct {
				Value int `json:"value"`
			} `json:"jvm.buffers.mapped.count"`
			JvmBuffersMappedUsed struct {
				Value int `json:"value"`
			} `json:"jvm.buffers.mapped.used"`
			JvmClassesLoaded struct {
				Value int `json:"value"`
			} `json:"jvm.classes.loaded"`
			JvmClassesUnloaded struct {
				Value int `json:"value"`
			} `json:"jvm.classes.unloaded"`
			JvmGcConcurrentMarkSweepCount struct {
				Value int `json:"value"`
			} `json:"jvm.gc.ConcurrentMarkSweep.count"`
			JvmGcConcurrentMarkSweepTime struct {
				Value int `json:"value"`
			} `json:"jvm.gc.ConcurrentMarkSweep.time"`
			JvmGcParNewCount struct {
				Value int `json:"value"`
			} `json:"jvm.gc.ParNew.count"`
			JvmGcParNewTime struct {
				Value int `json:"value"`
			} `json:"jvm.gc.ParNew.time"`
			JvmMemoryHeapCommitted struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.heap.committed"`
			JvmMemoryHeapInit struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.heap.init"`
			JvmMemoryHeapMax struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.heap.max"`
			JvmMemoryHeapUsage struct {
				Value float64 `json:"value"`
			} `json:"jvm.memory.heap.usage"`
			JvmMemoryHeapUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.heap.used"`
			JvmMemoryNonHeapCommitted struct {
				Value int `json:"value"`
			} `json:"jvm.memory.non-heap.committed"`
			JvmMemoryNonHeapInit struct {
				Value int `json:"value"`
			} `json:"jvm.memory.non-heap.init"`
			JvmMemoryNonHeapMax struct {
				Value int `json:"value"`
			} `json:"jvm.memory.non-heap.max"`
			JvmMemoryNonHeapUsage struct {
				Value int `json:"value"`
			} `json:"jvm.memory.non-heap.usage"`
			JvmMemoryNonHeapUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.non-heap.used"`
			JvmMemoryPoolsCMSOldGenCommitted struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.pools.CMS-Old-Gen.committed"`
			JvmMemoryPoolsCMSOldGenInit struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.pools.CMS-Old-Gen.init"`
			JvmMemoryPoolsCMSOldGenMax struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.pools.CMS-Old-Gen.max"`
			JvmMemoryPoolsCMSOldGenUsage struct {
				Value float64 `json:"value"`
			} `json:"jvm.memory.pools.CMS-Old-Gen.usage"`
			JvmMemoryPoolsCMSOldGenUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.CMS-Old-Gen.used"`
			JvmMemoryPoolsCodeCacheCommitted struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Code-Cache.committed"`
			JvmMemoryPoolsCodeCacheInit struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Code-Cache.init"`
			JvmMemoryPoolsCodeCacheMax struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Code-Cache.max"`
			JvmMemoryPoolsCodeCacheUsage struct {
				Value float64 `json:"value"`
			} `json:"jvm.memory.pools.Code-Cache.usage"`
			JvmMemoryPoolsCodeCacheUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Code-Cache.used"`
			JvmMemoryPoolsCompressedClassSpaceCommitted struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Compressed-Class-Space.committed"`
			JvmMemoryPoolsCompressedClassSpaceInit struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Compressed-Class-Space.init"`
			JvmMemoryPoolsCompressedClassSpaceMax struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Compressed-Class-Space.max"`
			JvmMemoryPoolsCompressedClassSpaceUsage struct {
				Value float64 `json:"value"`
			} `json:"jvm.memory.pools.Compressed-Class-Space.usage"`
			JvmMemoryPoolsCompressedClassSpaceUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Compressed-Class-Space.used"`
			JvmMemoryPoolsMetaspaceCommitted struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Metaspace.committed"`
			JvmMemoryPoolsMetaspaceInit struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Metaspace.init"`
			JvmMemoryPoolsMetaspaceMax struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Metaspace.max"`
			JvmMemoryPoolsMetaspaceUsage struct {
				Value float64 `json:"value"`
			} `json:"jvm.memory.pools.Metaspace.usage"`
			JvmMemoryPoolsMetaspaceUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Metaspace.used"`
			JvmMemoryPoolsParEdenSpaceCommitted struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Eden-Space.committed"`
			JvmMemoryPoolsParEdenSpaceInit struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Eden-Space.init"`
			JvmMemoryPoolsParEdenSpaceMax struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Eden-Space.max"`
			JvmMemoryPoolsParEdenSpaceUsage struct {
				Value float64 `json:"value"`
			} `json:"jvm.memory.pools.Par-Eden-Space.usage"`
			JvmMemoryPoolsParEdenSpaceUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Eden-Space.used"`
			JvmMemoryPoolsParSurvivorSpaceCommitted struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Survivor-Space.committed"`
			JvmMemoryPoolsParSurvivorSpaceInit struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Survivor-Space.init"`
			JvmMemoryPoolsParSurvivorSpaceMax struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Survivor-Space.max"`
			JvmMemoryPoolsParSurvivorSpaceUsage struct {
				Value float64 `json:"value"`
			} `json:"jvm.memory.pools.Par-Survivor-Space.usage"`
			JvmMemoryPoolsParSurvivorSpaceUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.pools.Par-Survivor-Space.used"`
			JvmMemoryTotalCommitted struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.total.committed"`
			JvmMemoryTotalInit struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.total.init"`
			JvmMemoryTotalMax struct {
				Value int64 `json:"value"`
			} `json:"jvm.memory.total.max"`
			JvmMemoryTotalUsed struct {
				Value int `json:"value"`
			} `json:"jvm.memory.total.used"`
			JvmThreadsBlockedCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.blocked.count"`
			JvmThreadsCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.count"`
			JvmThreadsDaemonCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.daemon.count"`
			JvmThreadsDeadlockCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.deadlock.count"`
			JvmThreadsDeadlocks struct {
				Value []interface{} `json:"value"`
			} `json:"jvm.threads.deadlocks"`
			JvmThreadsNewCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.new.count"`
			JvmThreadsRunnableCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.runnable.count"`
			JvmThreadsTerminatedCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.terminated.count"`
			JvmThreadsTimedWaitingCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.timed_waiting.count"`
			JvmThreadsWaitingCount struct {
				Value int `json:"value"`
			} `json:"jvm.threads.waiting.count"`
			ComDuitangBuyCartNapiCartAPIAddCartTime struct {
				Value float64 `json:"value"`
			} `json:"com.duitang.buy.cart.napi.CartAPI.addCartTime"`
			ComDuitangBuyDomainRepoSPUBrandRepoCacheHit struct {
				Value float64 `json:"value"`
			} `json:"com.duitang.buy.domain.repo.SPUBrandRepo.cacheHit"`
			ComDuitangBuyDomainRepoSPUCategoryRepoCacheHit struct {
				Value float64 `json:"value"`
			} `json:"com.duitang.buy.domain.repo.SPUCategoryRepo.cacheHit"`
			ComDuitangBuyDomainRepoSPUPropertiesRepoCacheHit struct {
				Value float64 `json:"value"`
			} `json:"com.duitang.buy.domain.repo.SPUPropertiesRepo.cacheHit"`
			ComDuitangBuyInventoryNapiInventoryAPIListCacheHit struct {
				Value float64 `json:"value"`
			} `json:"com.duitang.buy.inventory.napi.InventoryAPI.listCacheHit"`
			ComDuitangBuyViewNapiOrderAPICreateOrderTime struct {
				Value float64 `json:"value"`
			} `json:"com.duitang.buy.view.napi.OrderAPI.createOrderTime"`
			ComDuitangBuyViewNapiOrderListAPIListOrderTime struct {
				Value float64 `json:"value"`
			} `json:"com.duitang.buy.view.napi.OrderListAPI.listOrderTime"`
		} `json:"gauges"`
		Meters struct {
			AddCartAPI struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"add-cart-api"`
			AddCartAPIErr struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"add-cart-api-err"`
			ComDuitangBuyViewNapiPayPayAPIAlipay struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"com.duitang.buy.view.napi.pay.PayApi.-alipay"`
			ComDuitangBuyViewNapiPayPayAPIAlipayErr struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"com.duitang.buy.view.napi.pay.PayApi.-alipay-err"`
			ComDuitangBuyViewNapiPayPayAPIWechat struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"com.duitang.buy.view.napi.pay.PayApi.-wechat"`
			ComDuitangBuyViewNapiPayPayAPIWechatErr struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"com.duitang.buy.view.napi.pay.PayApi.-wechat-err"`
			CreateOrderAPI struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"create-order-api"`
			CreateOrderAPIErr struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"create-order-api-err"`
			OrderListAPI struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"order-list-api"`
			OrderListAPIErr struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"order-list-api-err"`
			WishListAPI struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"wish-list-api"`
			WishListAPIErr struct {
				Count    int     `json:"count"`
				M15Rate  float64 `json:"m15_rate"`
				M1Rate   float64 `json:"m1_rate"`
				M5Rate   float64 `json:"m5_rate"`
				MeanRate float64 `json:"mean_rate"`
				Units    string  `json:"units"`
			} `json:"wish-list-api-err"`
		} `json:"meters"`
		Timers struct {
			AddCartAPITimer struct {
				Count         int     `json:"count"`
				Max           float64 `json:"max"`
				Mean          float64 `json:"mean"`
				Min           float64 `json:"min"`
				P50           float64 `json:"p50"`
				P75           float64 `json:"p75"`
				P95           float64 `json:"p95"`
				P98           float64 `json:"p98"`
				P99           float64 `json:"p99"`
				P999          float64 `json:"p999"`
				Stddev        float64 `json:"stddev"`
				M15Rate       float64 `json:"m15_rate"`
				M1Rate        float64 `json:"m1_rate"`
				M5Rate        float64 `json:"m5_rate"`
				MeanRate      float64 `json:"mean_rate"`
				DurationUnits string  `json:"duration_units"`
				RateUnits     string  `json:"rate_units"`
			} `json:"add-cart-api-timer"`
			ComDuitangBuyViewNapiPayPayAPIAlipayTimer struct {
				Count         int     `json:"count"`
				Max           float64 `json:"max"`
				Mean          float64 `json:"mean"`
				Min           float64 `json:"min"`
				P50           float64 `json:"p50"`
				P75           float64 `json:"p75"`
				P95           float64 `json:"p95"`
				P98           float64 `json:"p98"`
				P99           float64 `json:"p99"`
				P999          float64 `json:"p999"`
				Stddev        float64 `json:"stddev"`
				M15Rate       float64 `json:"m15_rate"`
				M1Rate        float64 `json:"m1_rate"`
				M5Rate        float64 `json:"m5_rate"`
				MeanRate      float64 `json:"mean_rate"`
				DurationUnits string  `json:"duration_units"`
				RateUnits     string  `json:"rate_units"`
			} `json:"com.duitang.buy.view.napi.pay.PayApi.-alipay-timer"`
			ComDuitangBuyViewNapiPayPayAPIWechatTimer struct {
				Count         int     `json:"count"`
				Max           float64 `json:"max"`
				Mean          float64 `json:"mean"`
				Min           float64 `json:"min"`
				P50           float64 `json:"p50"`
				P75           float64 `json:"p75"`
				P95           float64 `json:"p95"`
				P98           float64 `json:"p98"`
				P99           float64 `json:"p99"`
				P999          float64 `json:"p999"`
				Stddev        float64 `json:"stddev"`
				M15Rate       float64 `json:"m15_rate"`
				M1Rate        float64 `json:"m1_rate"`
				M5Rate        float64 `json:"m5_rate"`
				MeanRate      float64 `json:"mean_rate"`
				DurationUnits string  `json:"duration_units"`
				RateUnits     string  `json:"rate_units"`
			} `json:"com.duitang.buy.view.napi.pay.PayApi.-wechat-timer"`
			CreateOrderAPITimer struct {
				Count         int     `json:"count"`
				Max           float64 `json:"max"`
				Mean          float64 `json:"mean"`
				Min           float64 `json:"min"`
				P50           float64 `json:"p50"`
				P75           float64 `json:"p75"`
				P95           float64 `json:"p95"`
				P98           float64 `json:"p98"`
				P99           float64 `json:"p99"`
				P999          float64 `json:"p999"`
				Stddev        float64 `json:"stddev"`
				M15Rate       float64 `json:"m15_rate"`
				M1Rate        float64 `json:"m1_rate"`
				M5Rate        float64 `json:"m5_rate"`
				MeanRate      float64 `json:"mean_rate"`
				DurationUnits string  `json:"duration_units"`
				RateUnits     string  `json:"rate_units"`
			} `json:"create-order-api-timer"`
			OrderListAPITimer struct {
				Count         int     `json:"count"`
				Max           float64 `json:"max"`
				Mean          float64 `json:"mean"`
				Min           float64 `json:"min"`
				P50           float64 `json:"p50"`
				P75           float64 `json:"p75"`
				P95           float64 `json:"p95"`
				P98           float64 `json:"p98"`
				P99           float64 `json:"p99"`
				P999          float64 `json:"p999"`
				Stddev        float64 `json:"stddev"`
				M15Rate       float64 `json:"m15_rate"`
				M1Rate        float64 `json:"m1_rate"`
				M5Rate        float64 `json:"m5_rate"`
				MeanRate      float64 `json:"mean_rate"`
				DurationUnits string  `json:"duration_units"`
				RateUnits     string  `json:"rate_units"`
			} `json:"order-list-api-timer"`
			WishListAPITimer struct {
				Count         int     `json:"count"`
				Max           float64 `json:"max"`
				Mean          float64 `json:"mean"`
				Min           float64 `json:"min"`
				P50           float64 `json:"p50"`
				P75           float64 `json:"p75"`
				P95           float64 `json:"p95"`
				P98           float64 `json:"p98"`
				P99           float64 `json:"p99"`
				P999          float64 `json:"p999"`
				Stddev        float64 `json:"stddev"`
				M15Rate       float64 `json:"m15_rate"`
				M1Rate        float64 `json:"m1_rate"`
				M5Rate        float64 `json:"m5_rate"`
				MeanRate      float64 `json:"mean_rate"`
				DurationUnits string  `json:"duration_units"`
				RateUnits     string  `json:"rate_units"`
			} `json:"wish-list-api-timer"`
		} `json:"timers"`
	} `json:"metrics"`
}

func (s *Buy) toCollectData(dataChan *chan *CollectData) {
	var data CollectData
	data.Name = "buy"
	data.Data = make(map[string]interface{})
	data.Tags = make(map[string]string)
	x := s.Timestamp / 1000
	y := s.Timestamp - x*1000
	data.T = time.Unix(x, y)
	data.Tags["host"] = s.Instance.Host
	data.Data["JvmBuffersDirectCapacity"] = float64(s.Metrics.Gauges.JvmBuffersDirectCapacity.Value)
	data.Data["JvmBuffersDirectCount"] = float64(s.Metrics.Gauges.JvmBuffersDirectCount.Value)
	data.Data["JvmBuffersDirectUsed"] = float64(s.Metrics.Gauges.JvmBuffersDirectUsed.Value)
	data.Data["JvmBuffersMappedCapacity"] = float64(s.Metrics.Gauges.JvmBuffersMappedCapacity.Value)
	data.Data["JvmBuffersMappedCount"] = float64(s.Metrics.Gauges.JvmBuffersMappedCount.Value)
	data.Data["JvmBuffersMappedUsed"] = float64(s.Metrics.Gauges.JvmBuffersMappedUsed.Value)
	data.Data["JvmClassesLoaded"] = float64(s.Metrics.Gauges.JvmClassesLoaded.Value)
	data.Data["JvmClassesUnloaded"] = float64(s.Metrics.Gauges.JvmClassesUnloaded.Value)
	data.Data["JvmGcConcurrentMarkSweepCount"] = float64(s.Metrics.Gauges.JvmGcConcurrentMarkSweepCount.Value)
	data.Data["JvmGcConcurrentMarkSweepTime"] = float64(s.Metrics.Gauges.JvmGcConcurrentMarkSweepTime.Value)
	data.Data["JvmGcParNewCount"] = float64(s.Metrics.Gauges.JvmGcParNewCount.Value)
	data.Data["JvmGcParNewTime"] = float64(s.Metrics.Gauges.JvmGcParNewTime.Value)
	data.Data["JvmMemoryHeapCommitted"] = float64(s.Metrics.Gauges.JvmMemoryHeapCommitted.Value)
	data.Data["JvmMemoryHeapInit"] = float64(s.Metrics.Gauges.JvmMemoryHeapInit.Value)
	data.Data["JvmMemoryHeapMax"] = float64(s.Metrics.Gauges.JvmMemoryHeapMax.Value)
	data.Data["JvmMemoryHeapUsage"] = float64(s.Metrics.Gauges.JvmMemoryHeapUsage.Value)
	data.Data["JvmMemoryHeapUsed"] = float64(s.Metrics.Gauges.JvmMemoryHeapUsed.Value)
	data.Data["JvmMemoryNonHeapCommitted"] = float64(s.Metrics.Gauges.JvmMemoryNonHeapCommitted.Value)
	data.Data["JvmMemoryNonHeapInit"] = float64(s.Metrics.Gauges.JvmMemoryNonHeapInit.Value)
	data.Data["JvmMemoryNonHeapMax"] = float64(s.Metrics.Gauges.JvmMemoryNonHeapMax.Value)
	data.Data["JvmMemoryNonHeapUsage"] = float64(s.Metrics.Gauges.JvmMemoryNonHeapUsage.Value)
	data.Data["JvmMemoryNonHeapUsed"] = float64(s.Metrics.Gauges.JvmMemoryNonHeapUsed.Value)
	data.Data["JvmMemoryPoolsCMSOldGenCommitted"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCMSOldGenCommitted.Value)
	data.Data["JvmMemoryPoolsCMSOldGenInit"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCMSOldGenInit.Value)
	data.Data["JvmMemoryPoolsCMSOldGenMax"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCMSOldGenMax.Value)
	data.Data["JvmMemoryPoolsCMSOldGenUsage"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCMSOldGenUsage.Value)
	data.Data["JvmMemoryPoolsCMSOldGenUsed"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCMSOldGenUsed.Value)
	data.Data["JvmMemoryPoolsCodeCacheCommitted"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCodeCacheCommitted.Value)
	data.Data["JvmMemoryPoolsCodeCacheInit"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCodeCacheInit.Value)
	data.Data["JvmMemoryPoolsCodeCacheMax"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCodeCacheMax.Value)
	data.Data["JvmMemoryPoolsCodeCacheUsage"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCodeCacheUsage.Value)
	data.Data["JvmMemoryPoolsCodeCacheUsed"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCodeCacheUsed.Value)
	data.Data["JvmMemoryPoolsCompressedClassSpaceCommitted"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCompressedClassSpaceCommitted.Value)
	data.Data["JvmMemoryPoolsCompressedClassSpaceInit"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCompressedClassSpaceInit.Value)
	data.Data["JvmMemoryPoolsCompressedClassSpaceMax"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCompressedClassSpaceMax.Value)
	data.Data["JvmMemoryPoolsCompressedClassSpaceUsage"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCompressedClassSpaceUsage.Value)
	data.Data["JvmMemoryPoolsCompressedClassSpaceUsed"] = float64(s.Metrics.Gauges.JvmMemoryPoolsCompressedClassSpaceUsed.Value)
	data.Data["JvmMemoryPoolsMetaspaceCommitted"] = float64(s.Metrics.Gauges.JvmMemoryPoolsMetaspaceCommitted.Value)
	data.Data["JvmMemoryPoolsMetaspaceInit"] = float64(s.Metrics.Gauges.JvmMemoryPoolsMetaspaceInit.Value)
	data.Data["JvmMemoryPoolsMetaspaceMax"] = float64(s.Metrics.Gauges.JvmMemoryPoolsMetaspaceMax.Value)
	data.Data["JvmMemoryPoolsMetaspaceUsage"] = float64(s.Metrics.Gauges.JvmMemoryPoolsMetaspaceUsage.Value)
	data.Data["JvmMemoryPoolsMetaspaceUsed"] = float64(s.Metrics.Gauges.JvmMemoryPoolsMetaspaceUsed.Value)
	data.Data["JvmMemoryPoolsParEdenSpaceCommitted"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParEdenSpaceCommitted.Value)
	data.Data["JvmMemoryPoolsParEdenSpaceInit"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParEdenSpaceInit.Value)
	data.Data["JvmMemoryPoolsParEdenSpaceMax"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParEdenSpaceMax.Value)
	data.Data["JvmMemoryPoolsParEdenSpaceUsage"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParEdenSpaceUsage.Value)
	data.Data["JvmMemoryPoolsParEdenSpaceUsed"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParEdenSpaceUsed.Value)
	data.Data["JvmMemoryPoolsParSurvivorSpaceCommitted"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParSurvivorSpaceCommitted.Value)
	data.Data["JvmMemoryPoolsParSurvivorSpaceInit"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParSurvivorSpaceInit.Value)
	data.Data["JvmMemoryPoolsParSurvivorSpaceMax"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParSurvivorSpaceMax.Value)
	data.Data["JvmMemoryPoolsParSurvivorSpaceUsage"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParSurvivorSpaceUsage.Value)
	data.Data["JvmMemoryPoolsParSurvivorSpaceUsed"] = float64(s.Metrics.Gauges.JvmMemoryPoolsParSurvivorSpaceUsed.Value)
	data.Data["JvmMemoryTotalCommitted"] = float64(s.Metrics.Gauges.JvmMemoryTotalCommitted.Value)
	data.Data["JvmMemoryTotalInit"] = float64(s.Metrics.Gauges.JvmMemoryTotalInit.Value)
	data.Data["JvmMemoryTotalMax"] = float64(s.Metrics.Gauges.JvmMemoryTotalMax.Value)
	data.Data["JvmMemoryTotalUsed"] = float64(s.Metrics.Gauges.JvmMemoryTotalUsed.Value)
	data.Data["JvmThreadsBlockedCount"] = float64(s.Metrics.Gauges.JvmThreadsBlockedCount.Value)
	data.Data["JvmThreadsCount"] = float64(s.Metrics.Gauges.JvmThreadsCount.Value)
	data.Data["JvmThreadsDaemonCount"] = float64(s.Metrics.Gauges.JvmThreadsDaemonCount.Value)
	data.Data["JvmThreadsDeadlockCount"] = float64(s.Metrics.Gauges.JvmThreadsDeadlockCount.Value)
	data.Data["JvmThreadsNewCount"] = float64(s.Metrics.Gauges.JvmThreadsNewCount.Value)
	data.Data["JvmThreadsRunnableCount"] = float64(s.Metrics.Gauges.JvmThreadsRunnableCount.Value)
	data.Data["JvmThreadsTerminatedCount"] = float64(s.Metrics.Gauges.JvmThreadsTerminatedCount.Value)
	data.Data["JvmThreadsTimedWaitingCount"] = float64(s.Metrics.Gauges.JvmThreadsTimedWaitingCount.Value)
	data.Data["JvmThreadsWaitingCount"] = float64(s.Metrics.Gauges.JvmThreadsWaitingCount.Value)
	data.Data["ComDuitangBuyCartNapiCartAPIAddCartTime"] = float64(s.Metrics.Gauges.ComDuitangBuyCartNapiCartAPIAddCartTime.Value)
	data.Data["ComDuitangBuyDomainRepoSPUBrandRepoCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyDomainRepoSPUBrandRepoCacheHit.Value)
	data.Data["ComDuitangBuyDomainRepoSPUCategoryRepoCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyDomainRepoSPUCategoryRepoCacheHit.Value)
	data.Data["ComDuitangBuyDomainRepoSPUPropertiesRepoCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyDomainRepoSPUPropertiesRepoCacheHit.Value)
	data.Data["ComDuitangBuyInventoryNapiInventoryAPIListCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyInventoryNapiInventoryAPIListCacheHit.Value)
	data.Data["ComDuitangBuyViewNapiOrderAPICreateOrderTime"] = float64(s.Metrics.Gauges.ComDuitangBuyViewNapiOrderAPICreateOrderTime.Value)
	data.Data["ComDuitangBuyViewNapiOrderListAPIListOrderTime"] = float64(s.Metrics.Gauges.ComDuitangBuyViewNapiOrderListAPIListOrderTime.Value)

	data.Data["ComDuitangBuyCartNapiCartAPIAddCartTime"] = float64(s.Metrics.Gauges.ComDuitangBuyCartNapiCartAPIAddCartTime.Value)
	data.Data["ComDuitangBuyDomainRepoSPUBrandRepoCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyDomainRepoSPUBrandRepoCacheHit.Value)
	data.Data["ComDuitangBuyDomainRepoSPUCategoryRepoCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyDomainRepoSPUCategoryRepoCacheHit.Value)
	data.Data["ComDuitangBuyDomainRepoSPUPropertiesRepoCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyDomainRepoSPUPropertiesRepoCacheHit.Value)
	data.Data["ComDuitangBuyInventoryNapiInventoryAPIListCacheHit"] = float64(s.Metrics.Gauges.ComDuitangBuyInventoryNapiInventoryAPIListCacheHit.Value)
	data.Data["ComDuitangBuyViewNapiOrderAPICreateOrderTime"] = float64(s.Metrics.Gauges.ComDuitangBuyViewNapiOrderAPICreateOrderTime.Value)
	data.Data["ComDuitangBuyViewNapiOrderListAPIListOrderTime"] = float64(s.Metrics.Gauges.ComDuitangBuyViewNapiOrderListAPIListOrderTime.Value)

	data.Data["AddCartAPICount"] = float64(s.Metrics.Meters.AddCartAPI.Count)
	data.Data["AddCartAPIErrCount"] = float64(s.Metrics.Meters.AddCartAPIErr.Count)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayCount"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipay.Count)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayErrCount"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipayErr.Count)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatCount"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechat.Count)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatErrCount"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechatErr.Count)
	data.Data["CreateOrderAPICount"] = float64(s.Metrics.Meters.CreateOrderAPI.Count)
	data.Data["CreateOrderAPIErrCount"] = float64(s.Metrics.Meters.CreateOrderAPIErr.Count)
	data.Data["OrderListAPICount"] = float64(s.Metrics.Meters.OrderListAPI.Count)
	data.Data["OrderListAPIErrCount"] = float64(s.Metrics.Meters.OrderListAPIErr.Count)
	data.Data["WishListAPICount"] = float64(s.Metrics.Meters.WishListAPI.Count)
	data.Data["WishListAPIErrCount"] = float64(s.Metrics.Meters.WishListAPIErr.Count)

	data.Data["AddCartAPIM1Rate"] = float64(s.Metrics.Meters.AddCartAPI.M1Rate)
	data.Data["AddCartAPIErrM1Rate"] = float64(s.Metrics.Meters.AddCartAPIErr.M1Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayM1Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipay.M1Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayErrM1Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipayErr.M1Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatM1Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechat.M1Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatErrM1Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechatErr.M1Rate)
	data.Data["CreateOrderAPIM1Rate"] = float64(s.Metrics.Meters.CreateOrderAPI.M1Rate)
	data.Data["CreateOrderAPIErrM1Rate"] = float64(s.Metrics.Meters.CreateOrderAPIErr.M1Rate)
	data.Data["OrderListAPIM1Rate"] = float64(s.Metrics.Meters.OrderListAPI.M1Rate)
	data.Data["OrderListAPIErrM1Rate"] = float64(s.Metrics.Meters.OrderListAPIErr.M1Rate)
	data.Data["WishListAPIM1Rate"] = float64(s.Metrics.Meters.WishListAPI.M1Rate)
	data.Data["WishListAPIErrM1Rate"] = float64(s.Metrics.Meters.WishListAPIErr.M1Rate)

	data.Data["AddCartAPIM15Rate"] = float64(s.Metrics.Meters.AddCartAPI.M15Rate)
	data.Data["AddCartAPIErrM15Rate"] = float64(s.Metrics.Meters.AddCartAPIErr.M15Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayM15Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipay.M15Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayErrM15Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipayErr.M15Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatM15Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechat.M15Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatErrM15Rate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechatErr.M15Rate)
	data.Data["CreateOrderAPIM15Rate"] = float64(s.Metrics.Meters.CreateOrderAPI.M15Rate)
	data.Data["CreateOrderAPIErrM15Rate"] = float64(s.Metrics.Meters.CreateOrderAPIErr.M15Rate)
	data.Data["OrderListAPIM15Rate"] = float64(s.Metrics.Meters.OrderListAPI.M15Rate)
	data.Data["OrderListAPIErrM15Rate"] = float64(s.Metrics.Meters.OrderListAPIErr.M15Rate)
	data.Data["WishListAPIM15Rate"] = float64(s.Metrics.Meters.WishListAPI.M15Rate)
	data.Data["WishListAPIErrM15Rate"] = float64(s.Metrics.Meters.WishListAPIErr.M15Rate)

	data.Data["AddCartAPIMeanRate"] = float64(s.Metrics.Meters.AddCartAPI.MeanRate)
	data.Data["AddCartAPIErrMeanRate"] = float64(s.Metrics.Meters.AddCartAPIErr.MeanRate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayMeanRate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipay.MeanRate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayErrMeanRate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIAlipayErr.MeanRate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatMeanRate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechat.MeanRate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIWechatErrMeanRate"] = float64(s.Metrics.Meters.ComDuitangBuyViewNapiPayPayAPIWechatErr.MeanRate)
	data.Data["CreateOrderAPIMeanRate"] = float64(s.Metrics.Meters.CreateOrderAPI.MeanRate)
	data.Data["CreateOrderAPIErrMeanRate"] = float64(s.Metrics.Meters.CreateOrderAPIErr.MeanRate)
	data.Data["OrderListAPIMeanRate"] = float64(s.Metrics.Meters.OrderListAPI.MeanRate)
	data.Data["OrderListAPIErrMeanRate"] = float64(s.Metrics.Meters.OrderListAPIErr.MeanRate)
	data.Data["WishListAPIMeanRate"] = float64(s.Metrics.Meters.WishListAPI.MeanRate)
	data.Data["WishListAPIErrMeanRate"] = float64(s.Metrics.Meters.WishListAPIErr.MeanRate)

	data.Data["AddCartAPITimer"] = float64(s.Metrics.Timers.AddCartAPITimer)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayTimer"] = float64(s.Metrics.Timers.ComDuitangBuyViewNapiPayPayAPIAlipayTimer)
	data.Data["CreateOrderAPITimer"] = float64(s.Metrics.Timers.CreateOrderAPITimer)
	data.Data["OrderListAPITimer"] = float64(s.Metrics.Timers.OrderListAPITimer)
	data.Data["WishListAPITimer"] = float64(s.Metrics.Timers.WishListAPITimer)

	data.Data["AddCartAPITimerCount"] = float64(s.Metrics.Timers.AddCartAPITimer.Count)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayTimerCount"] = float64(s.Metrics.Timers.ComDuitangBuyViewNapiPayPayAPIAlipayTimer.Count)
	data.Data["CreateOrderAPITimerCount"] = float64(s.Metrics.Timers.CreateOrderAPITimer.Count)
	data.Data["OrderListAPITimerCount"] = float64(s.Metrics.Timers.OrderListAPITimer.Count)
	data.Data["WishListAPITimerCount"] = float64(s.Metrics.Timers.WishListAPITimer.Count)

	data.Data["AddCartAPITimerMean"] = float64(s.Metrics.Timers.AddCartAPITimer.Mean)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayTimerMean"] = float64(s.Metrics.Timers.ComDuitangBuyViewNapiPayPayAPIAlipayTimer.Mean)
	data.Data["CreateOrderAPITimerMean"] = float64(s.Metrics.Timers.CreateOrderAPITimer.Mean)
	data.Data["OrderListAPITimerMean"] = float64(s.Metrics.Timers.OrderListAPITimer.Mean)
	data.Data["WishListAPITimerMean"] = float64(s.Metrics.Timers.WishListAPITimer.Mean)

	data.Data["AddCartAPITimerP95"] = float64(s.Metrics.Timers.AddCartAPITimer.P95)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayTimerP95"] = float64(s.Metrics.Timers.ComDuitangBuyViewNapiPayPayAPIAlipayTimer.P95)
	data.Data["CreateOrderAPITimerP95"] = float64(s.Metrics.Timers.CreateOrderAPITimer.P95)
	data.Data["OrderListAPITimerP95"] = float64(s.Metrics.Timers.OrderListAPITimer.P95)
	data.Data["WishListAPITimerP95"] = float64(s.Metrics.Timers.WishListAPITimer.P95)

	data.Data["AddCartAPITimerP99"] = float64(s.Metrics.Timers.AddCartAPITimer.P99)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayTimerP99"] = float64(s.Metrics.Timers.ComDuitangBuyViewNapiPayPayAPIAlipayTimer.P99)
	data.Data["CreateOrderAPITimerP99"] = float64(s.Metrics.Timers.CreateOrderAPITimer.P99)
	data.Data["OrderListAPITimerP99"] = float64(s.Metrics.Timers.OrderListAPITimer.P99)
	data.Data["WishListAPITimerP99"] = float64(s.Metrics.Timers.WishListAPITimer.P99)

	data.Data["AddCartAPITimerM1Rate"] = float64(s.Metrics.Timers.AddCartAPITimer.M1Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayTimerM1Rate"] = float64(s.Metrics.Timers.ComDuitangBuyViewNapiPayPayAPIAlipayTimer.M1Rate)
	data.Data["CreateOrderAPITimerM1Rate"] = float64(s.Metrics.Timers.CreateOrderAPITimer.M1Rate)
	data.Data["OrderListAPITimerM1Rate"] = float64(s.Metrics.Timers.OrderListAPITimer.M1Rate)
	data.Data["WishListAPITimerM1Rate"] = float64(s.Metrics.Timers.WishListAPITimer.M1Rate)

	data.Data["AddCartAPITimerM15Rate"] = float64(s.Metrics.Timers.AddCartAPITimer.M15Rate)
	data.Data["ComDuitangBuyViewNapiPayPayAPIAlipayTimerM15Rate"] = float64(s.Metrics.Timers.ComDuitangBuyViewNapiPayPayAPIAlipayTimer.M15Rate)
	data.Data["CreateOrderAPITimerM15Rate"] = float64(s.Metrics.Timers.CreateOrderAPITimer.M15Rate)
	data.Data["OrderListAPITimerM15Rate"] = float64(s.Metrics.Timers.OrderListAPITimer.M15Rate)
	data.Data["WishListAPITimerM15Rate"] = float64(s.Metrics.Timers.WishListAPITimer.M15Rate)

	*dataChan <- &data
}
