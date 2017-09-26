package collectors

import (
	"time"
)

type Munich struct {
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
		} `json:"gauges"`
	} `json:"metrics"`
}

func (s *Munich) toCollectData(dataChan *chan *CollectData) {
	var data CollectData
	data.Name = "munich"
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
	*dataChan <- &data
}
