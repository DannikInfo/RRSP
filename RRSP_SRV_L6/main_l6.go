package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("localhost"), Port: 8080})
	for {
		handleClient(listener)
	}
}

func handleClient(conn *net.UDPConn) {
	buf := make([]byte, 128)

	readLen, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.WriteToUDP(append([]byte("Hello, you said: "), buf[:readLen]...), addr)
}
