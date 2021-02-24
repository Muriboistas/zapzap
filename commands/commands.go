package commands

import (
	"errors"
	"fmt"
	"strings"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"

	cfg "github.com/muriboistas/zapzap/config"
	message "github.com/muriboistas/zapzap/infra/whats/message"
	"github.com/muriboistas/zapzap/pkg/helper/slicex"
	"github.com/muriboistas/zapzap/pkg/helper/stringx"
)

var config = cfg.Get.Command

// ActiveCommands the current avaliable commands
var (
	ActiveCommands = make(map[string]*Command)
	WaitList       = make(map[string]time.Time)
)

// Command all command data
type Command struct {
	ID          string
	Name        string
	Aliases     []string
	Subcommands []string
	Args        []string
	Help        string

	RootOnly bool

	Cooldown time.Duration

	Exec func(*whatsapp.Conn, whatsapp.TextMessage, map[string]string) error
}

// New creates a new command
func New(name string, f func(*whatsapp.Conn, whatsapp.TextMessage, map[string]string) error) *Command {
	return &Command{
		Name: strings.ToLower(name),
		Exec: f,
	}
}

// SetSubcommands set a sub command
func (c *Command) SetSubcommands(commands ...string) *Command {
	c.Subcommands = commands
	return c
}

// SetArgs set the command args
func (c *Command) SetArgs(args ...string) *Command {
	c.Args = slicex.Unique(args)
	return c
}

// SetHelp to the command
func (c *Command) SetHelp(format string, args ...interface{}) *Command {
	c.Help = fmt.Sprintf(format, args...)
	return c
}

// SetCooldown to re use the command
func (c *Command) SetCooldown(seconds time.Duration) *Command {
	c.Cooldown = seconds * time.Second
	return c
}

// SetAliases to the command
func (c *Command) SetAliases(aliases ...string) *Command {
	c.Aliases = aliases
	return c
}

// OnlyRoot only the bot number can use the command
func (c *Command) OnlyRoot() *Command {
	c.RootOnly = true
	return c
}

// Add activate command
func (c *Command) Add() {
	c.ID = makeCommandID(c)
	if !slicex.FoundString(config.Deactivate, c.ID) {
		ActiveCommands[strings.ToLower(c.ID)] = c
	}
}

// HavePermitions Check if the participant have the permitions to use some command
func HavePermitions(command *Command, msg whatsapp.TextMessage) (logs []string) {
	if !isRoot(command, msg) {
		logs = append(logs, "👾: You do not have permition to use this!")
	}

	return
}

// GetCommandID slicing the message and searching for aliases
func GetCommandID(text string) string {
	msgList := strings.Split(strings.Fields(text)[0], config.Prefix)
	cmd := msgList[1:]

	commandID := strings.Join(cmd, "-")
out:
	for _, command := range ActiveCommands {
		if commandID == command.ID {
			break out
		}
		for _, alias := range command.Aliases {
			if len(cmd) == 1 && alias == cmd[0] {
				commandID = makeCommandID(command)
				break out
			}
		}
	}

	return commandID
}

func makeCommandID(command *Command) string {
	return strings.TrimSuffix(fmt.Sprintf("%s-%s", command.Name, strings.Join(command.Subcommands, "-")), "-")
}

func trimCommand(msg string, command *Command) string {
	msgList := strings.Split(msg, " ")
	msg = strings.Join(msgList[1:], " ")

	return msg
}

func getCommandArgs(msg string, command *Command) (map[string]string, error) {
	args := make(map[string]string)
	msgArgs := stringx.ToArgs(msg)

	if len(msgArgs) > len(command.Args) {
		return nil, errors.New("You put more arguments than necessary, check the command again")
	}

	for k, argName := range command.Args {
		if k < len(msgArgs) {
			args[argName] = msgArgs[k]
		} else {
			args[argName] = ""
		}
	}
	return args, nil
}

func checkCooldown(command *Command, msg whatsapp.TextMessage) error {
	// get the message sender number
	num := message.GetSenderNumber(msg)
	if num == "" {
		return errors.New("blank number")
	}

	// verify if has cooldown
	cooldownID := num + command.ID
	if cd, found := WaitList[cooldownID]; found && !time.Now().After(cd) {
		return errors.New("command in cooldown")
	}

	// if need cooldown add it
	if command.Cooldown != time.Duration(0) {
		WaitList[cooldownID] = time.Now().Add(command.Cooldown)
	}

	return nil
}

// ParseCommand analyze the command
// FIXME: add error definitions
func ParseCommand(wac *whatsapp.Conn, msg whatsapp.TextMessage) {
	commandID := GetCommandID(msg.Text)
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

	// check if the commandis in cooldown an if has cooldown, if have add it
	if err := checkCooldown(command, msg); err != nil {
		message.Reply(err.Error(), wac, msg)
	}

	// trim command from message
	argsStr := trimCommand(msg.Text, command)
	args, err := getCommandArgs(argsStr, command)
	if err != nil {
		message.Reply("👾: "+err.Error(), wac, msg)
		return
	}

	err = command.Exec(wac, msg, args)
	if err != nil {
		message.Reply("👾: "+err.Error(), wac, msg)
	}
}
