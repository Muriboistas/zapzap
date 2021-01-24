package commands

import (
	"fmt"
	"strings"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"

	"github.com/muriboistas/zapzap/config"
	"github.com/muriboistas/zapzap/infra/whats/message"
)

// ActiveCommands the current avaliable commands
var (
	ActiveCommands = make(map[string]Command)
	WaitList       = make(map[string]time.Time)
)

// Command all command data
type Command struct {
	Name string
	Help string

	RootOnly bool

	Cooldown time.Duration

	Exec func(*whatsapp.Conn, whatsapp.TextMessage) error
}

// ParseCommand analyze the command
// FIXME: add error definitions
func ParseCommand(wac *whatsapp.Conn, msg whatsapp.TextMessage) {
	config := config.Get.Command
	// split the message by spaces
	msgArgs := strings.Fields(msg.Text)
	if len(msgArgs) < 1 {
		return
	}
	commandName := strings.ToLower(strings.TrimPrefix(msgArgs[0], config.Prefix))

	// verify if command exists
	command, found := ActiveCommands[commandName]
	if !found || commandName != command.Name {
		return
	}

	// get the message sender number
	num := message.GetSenderNumber(msg)
	if num == "" {
		return
	}

	// verify if has cooldown
	cooldownID := num + command.Name
	if cd, found := WaitList[cooldownID]; found && !time.Now().After(cd) {
		return
	}
	// if need cooldown add it
	if command.Cooldown != time.Duration(0) {
		WaitList[cooldownID] = time.Now().Add(command.Cooldown)
	}

	// verify if the message sender have the permitions
	if !HavePermitions(command, msg) {
		return
	}

	// trim command from message
	msg.Text = strings.TrimPrefix(msg.Text, msgArgs[0]+" ")

	err := command.Exec(wac, msg)
	if err != nil {
		message.Reply(err.Error(), wac, msg)
	}
}

// HavePermitions Check if the participant have the permitions to use some command
func HavePermitions(command Command, msg whatsapp.TextMessage) bool {
	// if command is root only
	if command.RootOnly && !msg.Info.FromMe {
		return false
	}

	return true
}

// New creates a new command
func New(name string, f func(*whatsapp.Conn, whatsapp.TextMessage) error) Command {
	return Command{
		Name: name,
		Exec: f,
	}
}

// SetHelp to the command
func (c Command) SetHelp(format string, args ...interface{}) Command {
	c.Help = fmt.Sprintf(format, args...)
	return c
}

// SetCooldown to re use the command
func (c Command) SetCooldown(seconds time.Duration) Command {
	c.Cooldown = seconds * time.Second
	return c
}

// OnlyRoot only the bot number can use the command
func (c Command) OnlyRoot() Command {
	c.RootOnly = true
	return c
}

// Add activate command
func (c Command) Add() {
	ActiveCommands[strings.ToLower(c.Name)] = c
}
