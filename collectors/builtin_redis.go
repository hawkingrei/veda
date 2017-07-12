package collectors

import (
	"strconv"
	"strings"
	"time"

	rd "github.com/garyburd/redigo/redis"
)

//import (
//	//"fmt"
//)

type RedisConnection struct {
	conn rd.Pool
	Name string
	tag  map[string]string
}

var _ Collectd = &RedisConnection{}

var (
	redisWhitelist = []string{
		"lru_clock",
		"connected_clients",
		"client_longest_output_list",
		"client_biggest_input_buf",
		"blocked_clients",
		"used_memory",
		"used_memory_rss",
		"used_memory_peak",
		"used_memory_lua",
		"mem_fragmentation_ratio",
		"loading",
		"rdb_changes_since_last_save",
		"rdb_bgsave_in_progress",
		"rdb_last_save_time",
		"rdb_last_bgsave_status",
		"rdb_last_bgsave_time_sec",
		"rdb_current_bgsave_time_sec",
		"aof_enabled",
		"aof_rewrite_in_progress",
		"aof_rewrite_scheduled",
		"aof_last_rewrite_time_sec",
		"aof_current_rewrite_time_sec",
		"aof_current_size",
		"total_connections_received",
		"total_commands_processed",
		"instantaneous_ops_per_sec",
		"rejected_connections",
		"sync_full",
		"sync_partial_ok",
		"sync_partial_err",
		"expired_keys",
		"evicted_keys",
		"keyspace_hits",
		"keyspace_misses",
		"pubsub_channels",
		"pubsub_patterns",
		"latest_fork_usec",
		"connected_slaves",
		"master_repl_offset",
		"repl_backlog_active",
		"repl_backlog_size",
		"repl_backlog_first_byte_offset",
		"repl_backlog_histlen",
		"used_cpu_sys",
		"used_cpu_user",
		"used_cpu_sys_children",
		"used_cpu_user_children",
	}
)

func GetRedisConn(address string, name string) (conn *RedisConnection, err error) {
	var tag map[string]string
	tag = make(map[string]string)
	tag["name"] = name

	nc := rd.Pool{
		MaxIdle:     1,
		IdleTimeout: 3 * time.Second,
		Dial: func() (rd.Conn, error) {
			c, err := rd.Dial("tcp", address,
				rd.DialConnectTimeout(3*time.Second),
				rd.DialReadTimeout(10*time.Second),
				rd.DialWriteTimeout(10*time.Second))
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	return &RedisConnection{
		conn: nc,
		Name: "redis",
		tag:  tag,
	}, err
}

func (mc *RedisConnection) convertCollectData(info string) (values CollectData) {
	values.T = time.Now()
	values.Name = mc.Name
	values.Tags = mc.tag
	values.Data = make(map[string]interface{})
	// Feed every line into
	result := strings.Split(info, "\r\n")
	// Load redis info values into array
	for _, value := range result {
		// Values are separated by :
		parts := ParseRedisLine(value, ":")
		if len(parts) == 2 {
			//fmt.Println(parts)
			itemValue, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				//fmt.Println("continue")
				continue
				// TODO: add log
			}
			for _, v := range redisWhitelist {
				if v == parts[0] {
					//fmt.Println("ok")
					values.Data[parts[0]] = itemValue
					break
				}
			}
		}
	}
	return values
}

func ParseRedisLine(s string, delimeter string) []string {
	return strings.Split(s, delimeter)
}

// Start a Redis Connection.
func (mc *RedisConnection) Start() (data CollectData, err error) {
	connect := mc.conn.Get()
	out, err := rd.String(connect.Do("INFO"))
	if err != nil {
		return data, err
	}
	data = mc.convertCollectData(out)
	return data, nil
}
