package collectors

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"
)

type BeansdbConnection struct {
	conn     net.Conn
	buffered bufio.ReadWriter
	Name     string
	tag      map[string]string
}

var (
	BeansdbWhitelist = []string{
		"uptime",
		"time",
		"pointer_size",
		"rusage_user",
		"rusage_system",
		"rusage_minflt",
		"rusage_majflt",
		"rusage_nswap",
		"rusage_inblock",
		"rusage_oublock",
		"rusage_nvcsw",
		"rusage_nivcsw",
		"item_buf_size",
		"curr_connections",
		"total_connections",
		"connection_structures",
		"cmd_get",
		"cmd_set",
		"cmd_delete",
		"slow_cmd",
		"get_hits",
		"get_misses",
		"curr_items",
		"total_items",
		"bytes_read",
		"bytes_written",
		"threads",
	}
)

func GetBeansdbConn(address string, name string) (conn *BeansdbConnection, err error) {
	var tag map[string]string
	tag = make(map[string]string)
	tag["name"] = name
	nc, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return conn, err
	}
	return &BeansdbConnection{
		conn: nc,
		buffered: bufio.ReadWriter{
			Reader: bufio.NewReader(nc),
			Writer: bufio.NewWriter(nc),
		},
		Name: "Beansdb",
		tag:  tag,
	}, err
}

func (bc *BeansdbConnection) writestring(s string) (err error) {
	_, err = bc.buffered.WriteString(s)
	return
}

func (bc *BeansdbConnection) flush() (err error) {
	return bc.buffered.Flush()
}

func (bc *BeansdbConnection) readline() string {
	bc.flush()
	l, _, _ := bc.buffered.ReadLine()
	return string(l)
}

func (bc *BeansdbConnection) stats() (result []byte, err error) {
	bc.writestring("stats\r\n")
	bc.flush()
	for {
		l := bc.readline()
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

func (bc *BeansdbConnection) convertCollectData(data string) (cd CollectData) {
	cd.T = time.Now()
	cd.Name = bc.Name
	cd.Tags = bc.tag
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
		for _, v := range BeansdbWhitelist {
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

// Start a Beansdb Connection.
func (bc *BeansdbConnection) Start() (data CollectData, err error) {
	result, err := bc.stats()
	if err != nil {
		return data, err
	}
	return bc.convertCollectData(string(result)), err
}
