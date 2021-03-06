package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	once sync.Once
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("client <ip> <port>")
		exit(0, nil)
	}

	defer func() {
		if err := recover(); err != nil {
			exit(1, err.(error))
		}
		exit(0, nil)
	}()

	// 1a. Request a connection to the server on the given IP/Port
	// 1b. Response from server should ask for credentials
	address := strings.Join(os.Args[1:], ":")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		exit(1, err)
	}

	fmt.Println("Connected to server at", address)

	defer conn.Close()

	go handleIncomingMessage(conn)
	handleStdIn(conn)

	// 2a. Client sends credentials
	// 2b. Response from server is either verified or not.
	// 2c. If verified, access to the server.
	// 2d. If not verified, drop the connection.

	// 3a. On connection, client can send messages to the server
}

func handleStdIn(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		conn.Write([]byte(scanner.Text() + "\n"))
	}

	if err := scanner.Err(); err != nil {
		exit(1, err)
	}
}

func handleIncomingMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			exit(1, err)
		}

		fmt.Print(message)
	}
}

func exit(code int, err error) {
	once.Do(func() {
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		os.Exit(code)
	})
}
