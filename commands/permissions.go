package commands

import (
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/muriboistas/zapzap/infra/whats/message"
)

func isRoot(command Command, msg whatsapp.TextMessage) bool {
	// if command is root only
	if command.RootOnly && !msg.Info.FromMe {
		return false
	}

	return true
}

func isInCooldown(command Command, msg whatsapp.TextMessage) bool {
	// get the message sender number
	num := message.GetSenderNumber(msg)
	if num == "" {
		return true
	}

	// verify if has cooldown
	cooldownID := num + command.Name
	if cd, found := WaitList[cooldownID]; found && !time.Now().After(cd) {
		return true
	}
	// if need cooldown add it
	if command.Cooldown != time.Duration(0) {
		WaitList[cooldownID] = time.Now().Add(command.Cooldown)
	}

	return false
}
