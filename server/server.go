package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/aisaacroth/chat-server/server/user"
)

var (
	once sync.Once
)

const (
	HOME = "127.0.0.1"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("server <port num>")
		exit(0, nil)
	}

	defer func() {
		if err := recover(); err != nil {
			exit(1, err.(error))
		}
		exit(0, nil)
	}()

	// 1a. Server reads in from a file of possible safe users?
	// 1b. Server can ask users to register with the service.
	address := strings.Join([]string{HOME, os.Args[1]}, ":")

	// 2a. Server listens on given port number for any incoming connections
	// 2b. Receives a request from client to access the system
	ln, err := net.Listen("tcp", address)
	if err != nil {
		exit(1, err)
	}

	defer ln.Close()

	fmt.Println("Server started at", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			exit(1, err)
		}

		fmt.Println("New connection from", conn.RemoteAddr().String())

		defer conn.Close()
		go handleConnection(conn)

	}
	// 3a. Send response requesting username and password
	// 3b. Receive response with username and password, verify against
	//     known users.
	// 3c. If the user does not validate self within 5 attempts, close
	//     the connection.
}

func handleConnection(conn net.Conn) {
	message := "Please input your username:"
	conn.Write([]byte(message + "\n"))
	reader := bufio.NewReader(conn)

	response, err := reader.ReadString('\n')

	if err != nil {
		exit(1, err)
	}

	newUser := user.User{response, conn.RemoteAddr()}
	message = "Welcome " + newUser.Name

	for {
		conn.Write([]byte(message + "\n"))
		response, err = bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Printf("%v has disconnected",
				conn.RemoteAddr().String())
			return
		}

		fmt.Printf("Response Received from %v: %v",
			conn.RemoteAddr().String(),
			string(response))
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
