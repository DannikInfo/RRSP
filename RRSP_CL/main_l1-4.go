package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	serv := os.Args[1]
	var conn net.Conn

	for conn == nil {
		conn, _ = net.Dial("tcp", serv)
		if conn == nil {
			fmt.Println("Connection error, try again after 5 sec...")
			time.Sleep(5 * time.Second)
		}
	}

	readFrom(conn)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	data, _ := reader.ReadString('\n')

	writeTo(conn, data)
	readFrom(conn)

}

func writeTo(dst net.Conn, data string) {
	if dst == nil {
		return
	}
	bufAnsw := make([]byte, 32)
	sended := false
	counter := 0
	for !sended && counter < 4 {
		counter++
		_, err := dst.Write([]byte(string(len(data))))
		if handleErr(err) {
			continue
		}
		readLen, err := dst.Read(bufAnsw)

		if string(bufAnsw[:readLen]) != "OK" || handleErr(err) {
			continue
		}

		_, err = dst.Write([]byte(data))
		if handleErr(err) {
			continue
		}

		sended = true
	}
}

func readFrom(conn net.Conn) {
	if conn == nil {
		return
	}
	bufLen := make([]byte, 32)
	readLen, err := conn.Read(bufLen)
	if handleErr(err) {
		return
	}

	varLen := int(bufLen[:readLen][0])

	_, err = conn.Write([]byte("OK"))

	buf := make([]byte, varLen)
	readLen, err = conn.Read(buf)

	if handleErr(err) {
		return
	}

	if string(buf[:readLen]) != "OK" {
		fmt.Print(string(buf[:readLen]))
	}
}

func handleErr(err error) bool {
	if err != nil {
		fmt.Printf("Error from TCP session: %s \n", err)
		return true
	}
	return false
}
