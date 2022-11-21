package main

import (
	"net"
	"strconv"
	"time"
)

func main() {
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("localhost"), Port: 8080})
	for {
		handleClient(listener)
	}
}

func handleClient(conn *net.UDPConn) {
	buf := make([]byte, 128)

	_, addr, _ := conn.ReadFromUDP(buf)
	for {
		conn.WriteToUDP([]byte(strconv.Itoa(int(time.Now().UnixMilli()))), addr)

		time.Sleep(200 * time.Millisecond)
	}
}
