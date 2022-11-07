package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
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

	writeTo(conn, "Hello, what's your name?\n")
	name := readFrom(conn)

	writeTo(conn, "Goodbye, "+name)
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

func readFrom(conn net.Conn) string {
	bufLen := make([]byte, 32)
	readLen, err := conn.Read(bufLen)
	if handleErr(err) {
		return ""
	}

	varLen := int(bufLen[:readLen][0])

	_, err = conn.Write([]byte("OK"))

	buf := make([]byte, varLen)
	readLen, err = conn.Read(buf)

	if handleErr(err) {
		return ""
	}

	if string(buf[:readLen]) != "OK" {
		return string(buf[:readLen])
	}

	return ""
}

func handleErr(err error) bool {
	if err != nil {
		fmt.Printf("Error from TCP session: %s \n", err)
		return true
	}
	return false
}
