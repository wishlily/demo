package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	ip := net.ParseIP("127.0.0.1")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	for i := 0; i < 5; i++ {
		conn.Write([]byte("hello"))

		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())

		buf.ReadBytes('\n')
	}
}
