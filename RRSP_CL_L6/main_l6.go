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
		conn, _ = net.Dial("udp", serv)
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

	for {
		conn.Write([]byte("give me time"))

		buf := make([]byte, 64)
		readLen, _ := conn.Read(buf)

		i, _ := strconv.ParseInt(string(buf[:readLen]), 10, 64)
		//tm := time.Unix(i, 0)
		fmt.Println(i) //, tm)
	}
}

func handleErr(err error) bool {
	if err != nil {
		fmt.Printf("Error from TCP session: %s \n", err)
		return true
	}
	return false
}
