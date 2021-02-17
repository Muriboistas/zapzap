package superflood

import (
	"errors"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New(
		"superflood", superflood,
	).SetArgs(
		"...",
	).SetHelp(
		"flood some message 150 or more times",
	).SetCooldown(1).OnlyRoot().Add()
}

func superflood(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	text := args["..."]
	if text == "" {
		return errors.New("You can not send blank messages")
	}
	message.Reply(text, wac, msg)
	for i := 0; i <= 150; i++ {
		message.Send(text, wac, msg)
	}

	return nil
}
