package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "[::1]:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		writeTo(conn, strconv.Itoa(int(time.Now().Unix())))
		time.Sleep(1 * time.Second)
	}
}

func writeTo(dst net.Conn, data string) {
	bufAnsw := make([]byte, 32)
	sended := false

	for !sended {
		_, err := dst.Write([]byte(string(len(data))))
		readLen, err := dst.Read(bufAnsw)

		if string(bufAnsw[:readLen]) != "OK" || handleErr(err) {
			continue
		}

		_, err = dst.Write([]byte(data))

		sended = true
	}
}

func handleErr(err error) bool {
	if err != nil {
		fmt.Printf("Error from TCP session: %s \n", err)
		return true
	}
	return false
}
