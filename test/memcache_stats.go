package main

// 写経 from: https://github.com/youtube/vitess/blob/master/go/memcache/memcache.go
//  * エラーハンドリングをできるかぎり省略して実装.
//  * statsコマンドだけ実装

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

type Connection struct {
	conn     net.Conn
	buffered bufio.ReadWriter
}

func GetConn(address string) (conn *Connection, err error) {
	nc, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Connection{
		conn: nc,
		buffered: bufio.ReadWriter{
			Reader: bufio.NewReader(nc),
			Writer: bufio.NewWriter(nc),
		},
	}, err
}

func (mc *Connection) writestring(s string) {
	_, err := mc.buffered.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func (mc *Connection) readline() string {
	mc.flush()
	l, _, _ := mc.buffered.ReadLine()
	return string(l)
}

func (mc *Connection) flush() {
	err := mc.buffered.Flush()
	if err != nil {
		panic(err)
	}
}

func (mc *Connection) Stats() (result []byte, err error) {
	mc.writestring("INFO")
	mc.flush()
	for {
		l := mc.readline()
		if strings.HasPrefix(l, "END") {
			break
		}
		if strings.Contains(l, "ERROR") {
			panic("ERROR")
		}
		result = append(result, l...)
		result = append(result, '\n')
	}
	return result, err
}

func main() {
	address := flag.String("address", "localhost:11211", "server address")
	flag.Parse()
	fmt.Printf("server: %s\n", *address)

	conn, err := GetConn(*address)
	if err != nil {
		fmt.Printf("%#v", err)
		os.Exit(1)
	}
	defer conn.conn.Close()

	result, _ := conn.Stats()
	fmt.Printf("%s", result)
}
