package help

import (
	"fmt"
	"sort"
	"strings"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/config"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

const helpHelp = "really?"

func init() {
	commands.New(
		"help", help,
	).SetAliases(
		"h",
	).SetArgs(
		"command",
	).SetHelp(
		helpHelp,
	).Add()
}

func help(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	res := getHelpDescription(msg, args)
	if res == "" {
		res = "invalid command"
	}
	message.Reply(res, wac, msg)

	return nil
}

func getHelpDescription(msg whatsapp.TextMessage, args map[string]string) string {
	command := args["command"]
	// if do not have args, list all avaliable commands
	if command == "" {
		commandList := make([]string, 0)
		for _, cmd := range commands.ActiveCommands {
			// verify if the user have permission to use the command
			if logs := commands.HavePermitions(cmd, msg); len(logs) == 0 {
				commandList = append(commandList, strings.ReplaceAll(cmd.ID, "-", config.Get.Command.Prefix))
			}
		}
		sort.Strings(commandList)
		return fmt.Sprintf("*Commands:*\n```%s```", strings.Join(commandList, "```\n```"))
	}

	cmdWithPref := command
	if !strings.HasPrefix(cmdWithPref, config.Get.Command.Prefix) {
		cmdWithPref = config.Get.Command.Prefix + command
	}
	commandID := commands.GetCommandID(cmdWithPref)
	if commandID == "" {
		return ""
	}

	// if argument is not blank
	if cmd, found := commands.ActiveCommands[commandID]; found {
		if cmd.Help == "" {
			cmd.Help = "no description"
		}

		originalCommand := fmt.Sprintf("%s%s", config.Get.Command.Prefix, strings.ReplaceAll(cmd.ID, "-", config.Get.Command.Prefix))
		helpMsg := fmt.Sprintf("*Command:* \n```%s```\n", originalCommand)
		// if command have args
		firstArg := "%s"
		if len(cmd.Args) > 0 {
			firstArg = "<%s>"
			helpMsg += "*Args:*\n```"
			helpMsg += strings.Join(cmd.Args, "```, ```")
			helpMsg += "```\n"
		}

		// if command have aliases
		if len(cmd.Aliases) > 0 {
			helpMsg += "*Aliases:*\n```" + config.Get.Command.Prefix
			helpMsg += strings.Join(cmd.Aliases, "```, ```"+config.Get.Command.Prefix)
			helpMsg += "```\n"
		}
		helpMsg += fmt.Sprintf("*Description:*\n_%s_", cmd.Help)

		examples := fmt.Sprintf("\n*Examples:*\n%s "+firstArg, originalCommand, strings.Join(cmd.Args, "> <"))
		for _, alias := range cmd.Aliases {
			examples += fmt.Sprintf("\n%s%s <%s>", config.Get.Command.Prefix, alias, strings.Join(cmd.Args, "> <"))
		}
		helpMsg += examples

		return helpMsg
	}

	return ""
}
