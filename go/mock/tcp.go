package main

import (
	"fmt"
	"net"
)

func get(conn net.Conn) (ret string, err error) {
	buf := make([]byte, 516)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return ret, err
	}
	fmt.Println(n)
	return string(n), nil
}
