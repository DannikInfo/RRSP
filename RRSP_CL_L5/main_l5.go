package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
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

	go readFrom(conn)

	exit := false
	for !exit {
		readFrom(conn)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		data, _ := reader.ReadString('\n')
		if strings.Compare(data, "exit") == 0 {
			exit = true
		}
	}
}

func readFrom(conn net.Conn) {
	count := 0
	for count < 4 {
		if conn == nil {
			count++
			continue
		}
		bufLen := make([]byte, 32)
		readLen, err := conn.Read(bufLen)
		if handleErr(err) {
			count++
			continue
		}

		varLen := int(bufLen[:readLen][0])

		_, err = conn.Write([]byte("OK"))

		buf := make([]byte, varLen)
		readLen, err = conn.Read(buf)

		if handleErr(err) {
			count++
			continue
		}

		if string(buf[:readLen]) != "OK" {
			i, _ := strconv.ParseInt(string(buf[:readLen]), 10, 64)
			tm := time.Unix(i, 0)
			fmt.Println(tm)
		}
		count = 0
	}
}

func handleErr(err error) bool {
	if err != nil {
		fmt.Printf("Error from TCP session: %s \n", err)
		return true
	}
	return false
}
