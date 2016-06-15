package user

import (
	"net"
)

type User struct {
	Name    string
	Address net.Addr
}

// String representation of user
func (u User) String() string {
	return "{" + string(u.Name) + "}"
}
