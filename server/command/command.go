package command

import (
	"strings"
)

type Command struct {
	SourceUser string
	DestUser   string
	Command    string
}

// String representation of a command
func (c Command) String() string {
	return "{" + strings.Join([]string{
		"Command:", c.Command, "From:", c.SourceUser,
		"To:", c.DestUser,
	}, " ") + "}"
}
