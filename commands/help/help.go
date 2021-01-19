package help

import (
	"fmt"
	"strings"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/config"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New("help", help).SetHelp("really?").Add()
}

func help(wac *whatsapp.Conn, msg whatsapp.TextMessage) {
	help := getHelpDescription(msg)
	message.Reply(help, wac, msg)
}

func getHelpDescription(msg whatsapp.TextMessage) string {
	command := config.Get.Command.Prefix + "help"
	text := strings.ToLower(msg.Text)
	// check if the message is just the help command
	if text != command {
		if cmd, found := commands.ActiveCommands[text]; found {
			if cmd.Help == "" {
				cmd.Help = "no description"
			}
			return fmt.Sprintf("*%s*: %s", cmd.Name, cmd.Help)
		}
		return "invalid command"
	}

	// list all commands
	commandList := make([]string, 0)
	for _, cmd := range commands.ActiveCommands {
		commandList = append(commandList, cmd.Name)
	}

	return fmt.Sprintf("*Commands:*\n```%s```", strings.Join(commandList, "```, ```"))

}
