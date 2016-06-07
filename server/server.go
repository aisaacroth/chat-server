package main

import (
    "fmt"
    "net"
    "os"
    "strings"
    "sync"
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

    address := strings.Join([]string{HOME, os.Args[1]}, "/")

    _, err := net.Listen("tcp", address)
    if err != nil {
    }

    fmt.Println("Server started at", address)


    // 1a. Server reads in from a file of possible safe users?
    // 1b. Server can ask users to register with the service.

    // 2a. Server listens on given port number for any incoming connections

    // 2b. Receives a request from client to access the system

    // 3a. Send response requesting username and password
    // 3b. Receive response with username and password, verify against
    //     known users.
    // 3c. If the user does not validate self within 5 attempts, close
    //     the connection.
}

func exit(code int, err error) {
    once.Do(func() {
        if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
        }

        os.Exit(code)
    })
}

