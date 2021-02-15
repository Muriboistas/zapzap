package commands

import (
	whatsapp "github.com/Rhymen/go-whatsapp"
)

func isRoot(command *Command, msg whatsapp.TextMessage) bool {
	// if command is root only
	if command.RootOnly && !msg.Info.FromMe {
		return false
	}

	return true
}
