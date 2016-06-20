package user

import (
	"net"
)

type User struct {
	Name     string
	Address  net.Addr
	Password string
}

// String representation of user
func (u User) String() string {
	return "{" + string(u.Name) + ":" + u.Address.String() + "}"
}
