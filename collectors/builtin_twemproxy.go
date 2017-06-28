package collectors

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"
)

type TwemproxyConnection struct {
	conn     net.Conn
	buffered bufio.ReadWriter
	Name     string
	tag      map[string]string
}

var _ Collectd = &TwemproxyConnection{}

var (
	memcacheWhitelist = []string{
		"uptime",
		"time",
		"pointer_size",
		"rusage_user",
		"rusage_user",
		"rusage_system",
		"curr_connections",
		"total_connections",
		"connection_structures",
		"cmd_get",
		"cmd_set",
		"cmd_flush",
		"get_hits",
		"get_misses",
		"delete_misses",
		"delete_hits",
		"incr_misses",
		"incr_hits",
		"decr_misses",
		"decr_hits",
		"cas_misses",
		"cas_hits",
		"cas_badval",
		"auth_cmds",
		"auth_errors",
		"bytes_read",
		"bytes_written",
		"limit_maxbytes",
		"accepting_conns",
		"listen_disabled_num",
		"threads",
		"conn_yields",
		"bytes",
		"curr_items",
		"total_items",
		"evictions",
		"reclaimed",
	}
)

func GetTwemproxyConn(address string, name string) (conn *MemcacheConnection, err error) {
	var tag map[string]string
	tag = make(map[string]string)
	tag["name"] = name
	nc, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return conn, err
	}
	return &MemcacheConnection{
		conn: nc,
		buffered: bufio.ReadWriter{
			Reader: bufio.NewReader(nc),
			Writer: bufio.NewWriter(nc),
		},
		Name: "memcache",
		tag:  tag,
	}, err
}

func (mc *TwemproxyConnection) writestring(s string) (err error) {
	_, err = mc.buffered.WriteString(s)
	return
}

func (mc *TwemproxyConnection) flush() (err error) {
	return mc.buffered.Flush()
}

func (mc *TwemproxyConnection) readline() string {
	mc.flush()
	l, _, _ := mc.buffered.ReadLine()
	return string(l)
}

func (mc *TwemproxyConnection) stats() (result []byte, err error) {
	mc.writestring("stats\r\n")
	mc.flush()
	for {
		l := mc.readline()
		if strings.HasPrefix(l, "END") {
			break
		}
		if strings.Contains(l, "ERROR") {
			return result, errors.New("GET ERROR WHEN GETTING STATS")
		}
		result = append(result, l...)
		result = append(result, '\n')
	}
	return result, err
}

func (mc *TwemproxyConnection) convertCollectData(data string) (cd CollectData) {
	cd.T = time.Now()
	cd.Name = mc.Name
	cd.Tags = mc.tag
	cd.Data = make(map[string]interface{})
	resultStr := strings.Split(data, "\n")
	for _, result := range resultStr {
		if result == "" {
			break
		}
		item := strings.Split(result, " ")
		itemValue, err := strconv.ParseFloat(item[2], 64)
		if err != nil {
			continue
			// TODO: add log
		}
		for _, v := range memcacheWhitelist {
			if v == item[1] {
				// Found!
				cd.Data[item[1]] = itemValue
				break
			}
		}
	}
	//fmt.Println(cd)
	return
}

// Start a Twemproxy Connection.
func (mc *TwemproxyConnection) Start() (data CollectData, err error) {
	result, err := mc.stats()
	if err != nil {
		return data, err
	}
	return mc.convertCollectData(string(result)), err
}
