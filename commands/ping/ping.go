package ping

import (
	"github.com/muriboistas/zapzap/commands"

	"github.com/muriboistas/zapzap/infra/whats/message"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New(
		"ping", ping,
	).SetHelp(
		"check the command",
	).SetCooldown(1).OnlyRoot().Add()
}

func ping(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	message.Reply("Pong", wac, msg)

	return nil
}
