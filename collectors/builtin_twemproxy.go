package collectors

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net"
	"time"
)

type TwemproxyConnection struct {
	conn     string
	buffered bufio.ReadWriter
	Name     string
	tag      map[string]string
}

var _ Collectd = &TwemproxyConnection{}

var (
	twemproxyWhitelist = []string{
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

type Session struct {
	Client_eof         float64 `json:"client_eof"`
	Client_err         float64 `json:"client_eff"`
	Client_connections float64 `json:"client_connections"`
	Server_ejects      float64 `json:"server_ejects"`
	Forward_error      float64 `json:"forward_error"`
	Fragments          float64 `json:"fragments"`
}

type twemproxyJson struct {
	Uptime  float64 `json:"uptime"`
	Session Session `json:"session"`
}

func GetTwemproxyConn(address string, name string) (conn *TwemproxyConnection, err error) {
	var tag map[string]string
	tag = make(map[string]string)
	tag["name"] = name

	return &TwemproxyConnection{
		conn: address,
		//buffered: bufio.ReadWriter{
		//	Reader: bufio.NewReader(nc),
		//	Writer: bufio.NewWriter(nc),
		//},
		Name: "memcache",
		tag:  tag,
	}, err
}

func (mc *TwemproxyConnection) stats() (result []byte, err error) {
	conn, err := net.DialTimeout("tcp", mc.conn, 2*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()
	return ioutil.ReadAll(conn)
}

func (mc *TwemproxyConnection) convertCollectData(data []byte) (cd CollectData) {
	cd.T = time.Now()
	cd.Name = mc.Name
	cd.Tags = mc.tag
	cd.Data = make(map[string]interface{})
	var da twemproxyJson
	json.Unmarshal(data, &da)
	//fmt.Println(da)
	cd.Data["uptime"] = da.Uptime
	cd.Data["client_connections"] = da.Session.Client_connections
	cd.Data["client_eof"] = da.Session.Client_eof
	cd.Data["client_err"] = da.Session.Client_err
	cd.Data["client_connections"] = da.Session.Client_connections
	cd.Data["server_ejects"] = da.Session.Server_ejects
	cd.Data["forward_error"] = da.Session.Forward_error
	cd.Data["fragments"] = da.Session.Fragments
	return
}

// Start a Twemproxy Connection.
func (mc *TwemproxyConnection) Start() (data CollectData, err error) {
	result, err := mc.stats()
	if err != nil {
		return data, err
	}
	//fmt.Println(string(result))
	return mc.convertCollectData(result), err
}
