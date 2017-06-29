package collectors

import (
	"bytes"
	"io/ioutil"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ZookeeperConnection struct {
	conn net.Conn
	Name string
	tag  map[string]string
}

var _ Collectd = &ZookeeperConnection{}

var (
	paramMatcher       = regexp.MustCompile("([^\\s]+)\\s+(.*$)")
	zookeeperWhitelist = []string{
		"zk_avg_latency",
		"zk_max_latency",
		"zk_min_latency",
		"zk_packets_received",
		"zk_packets_sent",
		"zk_num_alive_connections",
		"zk_outstanding_requests",
		"zk_znode_count",
		"zk_watch_count",
		"zk_ephemerals_count",
		"zk_approximate_data_size",
		"zk_open_file_descriptor_count",
		"zk_max_file_descriptor_count",
	}
)

func GetZookeeperConn(address string, name string) (conn *ZookeeperConnection, err error) {
	var tag map[string]string
	tag = make(map[string]string)
	tag["name"] = name
	nc, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return conn, err
	}

	return &ZookeeperConnection{
		conn: nc,
		Name: "zookeeper",
		tag:  tag,
	}, err
}

func (zc *ZookeeperConnection) stats() (result string, err error) {
	// Set read and write timeout.
	command := "mntr"
	// Write four-letter command.
	_, err = zc.conn.Write([]byte(command))
	if err != nil {
		return result, err
	}

	r, err := ioutil.ReadAll(zc.conn)
	if err != nil {
		return result, err
	}
	stream := bytes.NewReader(r)
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String(), nil
}

func (zc *ZookeeperConnection) convertCollectData(data string) (cd CollectData) {
	cd.T = time.Now()
	cd.Name = zc.Name
	cd.Tags = zc.tag
	cd.Data = make(map[string]interface{})
	resultStr := strings.Split(data, "\n")
	for _, result := range resultStr {
		if result == "" {
			break
		}
		//fmt.Println(item)
		if match := paramMatcher.FindStringSubmatch(result); len(match) == 3 {
			for _, v := range zookeeperWhitelist {
				if v == match[1] {
					itemValue, err := strconv.ParseFloat(match[2], 64)
					if err != nil {
						break
						// TODO: add log
					}
					// Found!
					cd.Data[match[1]] = itemValue
				}
			}
		}
	}
	return
}

func (zc *ZookeeperConnection) Start() (data CollectData, err error) {
	result, err := zc.stats()
	if err != nil {
		return data, err
	}
	//fmt.Println(result)
	return zc.convertCollectData(string(result)), err
}
