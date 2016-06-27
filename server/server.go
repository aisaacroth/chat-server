package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

        "github.com/aisaacroth/chat-server/server/command"
	"github.com/aisaacroth/chat-server/server/user"
)

var (
	once sync.Once
)

const (
	HOME = "127.0.0.1"
)

type UserList []user.User

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

	address := strings.Join([]string{HOME, os.Args[1]}, ":")

	ln, err := net.Listen("tcp", address)
	if err != nil {
		exit(1, err)
	}

	defer ln.Close()

	fmt.Println("Server started at", address)
        userSlice := make(UserList, 0)

	for {
		conn, err := ln.Accept()
		if err != nil {
			exit(1, err)
		}

		fmt.Println("New connection from", conn.RemoteAddr().String())

		defer conn.Close()
		go handleConnection(&userSlice, conn)

	}
}

func handleConnection(userSlice *UserList, conn net.Conn) {
	reader := bufio.NewReader(conn)
	newUser := RegisterNewUser(conn, reader)
        *userSlice = append(*userSlice, newUser)
	message := "Welcome " + newUser.Name

	fmt.Printf("New User created: %v\n", newUser)
        fmt.Printf("User Slice Size: %v\n", len(*userSlice))

	for {
		conn.Write([]byte(message + "\n"))
		response, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Printf("%v has disconnected\n",
				conn.RemoteAddr().String())
			return
		}

		fmt.Printf("Response Received from %v: %v",
			conn.RemoteAddr().String(),
			string(response))

                HandleCommand(message)
	}
}

func RegisterNewUser(conn net.Conn, reader *bufio.Reader) user.User {
	message := "Please input your username:"
	conn.Write([]byte(message + "\n"))

	response, err := reader.ReadString('\n')

	if err != nil {
		exit(1, err)
	}

	cleanName := strings.TrimSpace(response)

	newUser := user.User{cleanName, conn.RemoteAddr()}

	return newUser
}

func HandleCommand(command string,  conn net.Conn) {
    request := CreateCommand(command, conn)
}

func CreateCommand(command string, conn net.Conn) command.Command {
    return nil
}


func exit(code int, err error) {
	once.Do(func() {
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		os.Exit(code)
	})
}
