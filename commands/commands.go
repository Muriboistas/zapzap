package commands

import (
	"fmt"
	"strings"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"

	cfg "github.com/muriboistas/zapzap/config"
	message "github.com/muriboistas/zapzap/infra/whats/message"
)

var config = cfg.Get.Command

// ActiveCommands the current avaliable commands
var (
	ActiveCommands = make(map[string]Command)
	WaitList       = make(map[string]time.Time)
)

// Command all command data
type Command struct {
	ID          string
	Name        string
	Subcommands []string
	Args        []string
	Help        string

	RootOnly bool

	Cooldown time.Duration

	Exec func(*whatsapp.Conn, whatsapp.TextMessage, map[string]string) error
}

// New creates a new command
func New(name string, f func(*whatsapp.Conn, whatsapp.TextMessage, map[string]string) error) Command {
	return Command{
		Name: strings.ToLower(name),
		Exec: f,
	}
}

// SetSubcommand set a sub command
func (c Command) SetSubcommand(name string) Command {
	c.Subcommands = append(c.Subcommands, strings.ToLower(name))
	return c
}

// SetArgs set the command args
func (c Command) SetArgs(args ...string) Command {
	c.Args = args
	return c
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
	c.ID = strings.TrimSuffix(fmt.Sprintf("%s-%s", c.Name, strings.Join(c.Subcommands, "-")), "-")
	ActiveCommands[strings.ToLower(c.ID)] = c
}

// HavePermitions Check if the participant have the permitions to use some command
func HavePermitions(command Command, msg whatsapp.TextMessage) (logs []string) {
	if !isRoot(command, msg) {
		logs = append(logs, "👾: You do not have permition tu use this!")
	}

	if isInCooldown(command, msg) {
		logs = append(logs, "👾: Command in cooldown!")
	}

	return
}

func getCommandID(msg whatsapp.TextMessage) string {
	message := strings.ToLower(strings.TrimPrefix(msg.Text, config.Prefix))
	var commandID string
	var subCommandsLen int
	for k, v := range ActiveCommands {
		command := strings.TrimSpace(fmt.Sprintf("%s %s", v.Name, strings.Join(v.Subcommands, " ")))
		if !strings.HasPrefix(message, command) {
			continue
		}

		if subCommandsLen > len(v.Subcommands) {
			continue
		}

		subCommandsLen = len(v.Subcommands)
		commandID = k
	}

	return commandID
}

func trimCommand(msg string, command Command) string {
	msg = strings.TrimPrefix(msg, fmt.Sprintf("%s%s %s", config.Prefix, command.Name, strings.Join(command.Subcommands, " ")))
	msg = strings.TrimPrefix(msg, " ")

	return msg
}

func getCommandArgs(msg string, command Command) map[string]string {
	args := make(map[string]string)
	msgFields := strings.Fields(msg)

	for k, argName := range command.Args {
		// it will save the rest of fields in "..."
		if k < len(msgFields) {
			args[argName] = msgFields[k]
		} else {
			args[argName] = ""
		}
	}

	args["..."] = strings.Join(msgFields[len(command.Args):], " ")

	return args
}

// ParseCommand analyze the command
// FIXME: add error definitions
func ParseCommand(wac *whatsapp.Conn, msg whatsapp.TextMessage) {
	commandID := getCommandID(msg)
	if commandID == "" {
		return
	}

	// verify if command exists
	command, found := ActiveCommands[commandID]
	if !found || commandID != command.ID {
		return
	}

	// verify if the message sender have the permitions
	if logs := HavePermitions(command, msg); len(logs) > 0 {
		message.Reply(strings.Join(logs, "\n"), wac, msg)
		return
	}

	// trim command from message
	msg.Text = trimCommand(msg.Text, command)

	args := getCommandArgs(msg.Text, command)

	err := command.Exec(wac, msg, args)
	if err != nil {
		message.Reply("👾: "+err.Error(), wac, msg)
	}
}
