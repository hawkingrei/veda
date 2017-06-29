package collectors

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type NginxConnection struct {
	conn string
	Name string
	tag  map[string]string
}

var _ Collectd = &NginxConnection{}

var (
	activeRe  = regexp.MustCompile("Active connections: (\\d+)")
	requestRe = regexp.MustCompile("\\s(\\d+)\\s+(\\d+)\\s+(\\d+)")
	connRe    = regexp.MustCompile("Reading: (\\d+) Writing: (\\d+) Waiting: (\\d+)")
)

func GetNginxConn(address string, name string) (conn *NginxConnection, err error) {
	var tag map[string]string
	tag = make(map[string]string)
	tag["name"] = name
	return &NginxConnection{
		conn: address,
		Name: "nginx",
		tag:  tag,
	}, err
}

func (nc *NginxConnection) stats() (result string, err error) {
	resp, err := http.Get(nc.conn)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	return string(body), err
}

func (nc *NginxConnection) convertCollectData(data string) (cd CollectData) {
	cd.T = time.Now()
	cd.Name = nc.Name
	cd.Tags = nc.tag
	cd.Data = make(map[string]interface{})
	matches := activeRe.FindStringSubmatch(data)
	if matches == nil {
		return
	}
	active, _ := strconv.Atoi(matches[1])
	matches = requestRe.FindStringSubmatch(data)
	accepts, _ := strconv.Atoi(matches[1])
	handled, _ := strconv.Atoi(matches[2])
	requests, _ := strconv.Atoi(matches[3])
	dropped := accepts - handled
	matches = connRe.FindStringSubmatch(data)
	reading, _ := strconv.Atoi(matches[1])
	writing, _ := strconv.Atoi(matches[2])
	waiting, _ := strconv.Atoi(matches[3])
	cd.Data["active"] = active
	cd.Data["accepts"] = accepts
	cd.Data["handled"] = handled
	cd.Data["dropped"] = dropped
	cd.Data["requests"] = requests
	cd.Data["reading"] = reading
	cd.Data["writing"] = writing
	cd.Data["waiting"] = waiting
	return
}

func (nc *NginxConnection) Start() (data CollectData, err error) {
	result, err := nc.stats()
	if err != nil {
		return data, err
	}
	return nc.convertCollectData(string(result)), err
}
