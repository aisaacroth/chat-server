package user

import (
    "net"
)

type User struct {
    Name string
    Address net.Addr
}
