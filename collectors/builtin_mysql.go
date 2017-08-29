package collectors

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnection struct {
	DB   *sql.DB
	dsn  string
	Name string
	tag  map[string]string
}

var (
	mysqlBlacklist = []string{
		"Rpl_status",
		"Innodb_buffer_pool_load_status",
		"Ssl_session_cache_mode",
		"Innodb_buffer_pool_dump_status",
		"Slave_running",
		"Rpl_semi_sync_master_status",
		"Rpl_semi_sync_slave_status",
		"Innodb_have_atomic_builtins",
		"Ssl_version",
		"Sphinx_words",
		"Compression",
		"Binlog_snapshot_file",
	}
)

func GetMysqlConn(address string, username string, password string, dbname string, name string) (conn *MysqlConnection, err error) {
	var tag map[string]string
	tag = make(map[string]string)
	tag["name"] = name
	dsn := username + ":" + password + "@tcp(" + address + ")/" + dbname
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		return conn, err
	}
	return &MysqlConnection{
		DB:   DB,
		dsn:  dsn,
		Name: "mysql",
		tag:  tag,
	}, err
}

func (mc *MysqlConnection) loadStatus( /*db *sql.DB*/ ) (data map[string]string, err error) {
	// Returns the global status, also for versions previous 5.0.2
	rows, err := mc.DB.Query("SHOW /*!50002 GLOBAL */ STATUS;")
	if err != nil {
		return data, err
	}
	defer rows.Close()

	data = make(map[string]string)
WRITE:
	for rows.Next() {
		var name string
		var value string

		err = rows.Scan(&name, &value)
		if err != nil {
			return nil, err
		}
		for _, v := range mysqlBlacklist {
			if v == name {
				continue WRITE
			}
		}
		data[name] = value
	}

	return data, nil
}

func (mc *MysqlConnection) convertCollectData(info map[string]string) (values CollectData) {
	values.T = time.Now()
	values.Name = mc.Name
	values.Tags = mc.tag
	values.Data = make(map[string]interface{})
	for k, v := range info {
		value, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			values.Data[k] = value
		}
	}
	return
}

func (mc *MysqlConnection) Start() (data CollectData, err error) {
	d, err := mc.loadStatus()
	if err != nil {
		return data, err
	}
	return mc.convertCollectData(d), err
}
