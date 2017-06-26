package collectors

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"
)

type MemcacheConnection struct {
	conn     net.Conn
	buffered bufio.ReadWriter
	Name     string
	tag      map[string]string
	//writeChan
}

var _ Collectd = &MemcacheConnection{}

func GetMemcacheConn(address string, name string) (conn *MemcacheConnection, err error) {
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

func (mc *MemcacheConnection) writestring(s string) (err error) {
	_, err = mc.buffered.WriteString(s)
	return
}

func (mc *MemcacheConnection) flush() (err error) {
	return mc.buffered.Flush()
}

func (mc *MemcacheConnection) readline() string {
	mc.flush()
	l, _, _ := mc.buffered.ReadLine()
	return string(l)
}

func (mc *MemcacheConnection) stats() (result []byte, err error) {
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

func (mc *MemcacheConnection) convertCollectData(data []byte) (cd CollectData) {
	cd.T = time.Now()
	cd.Data = make(map[string]interface{})

	resultStr := strings.Split(string(data), "\n")
	for _, result := range resultStr {
		if result == "" {
			break
		}
		item := strings.Split(result, " ")
		itemValue, _ := strconv.ParseFloat(item[2], 64)
		cd.Data[item[1]] = itemValue
		//fmt.Println(cd.Data)
	}
	return
}

// Start a Memcache Connection.
func (mc *MemcacheConnection) Start() (data CollectData, err error) {
	result, err := mc.stats()
	if err != nil {
		return data, err
	}
	return mc.convertCollectData(result), err
}
