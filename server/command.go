package main

type commandId int

const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

// command type
type command struct {
	id commandID
	client *client	// this is the sender of the current command
	args []string	// arguments of the command
}
