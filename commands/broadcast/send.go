package broadcast

import (
	"errors"
	"time"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/config"
	bc "github.com/muriboistas/zapzap/infra/whats/broadcast"
	"github.com/muriboistas/zapzap/infra/whats/message"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New(
		"broadcast", sendBroadcast,
	).SetSubcommands(
		"send",
	).SetArgs(
		"...",
	).SetAliases(
		"bs",
	).SetHelp(
		"send message to all broadcasts in the active broadcast list",
	).SetCooldown(1).OnlyRoot().Add()
}

func sendBroadcast(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	text := args["..."]
	if text == "" {
		return errors.New("You can not send blank messages")
	}

	for _, remoteJid := range bc.Active {
		err := message.SendTo(remoteJid, text, wac)
		if err != nil {
			message.Reply("Failed while sending to "+remoteJid, wac, msg)
		} else {
			message.Reply("Successfully sendend to "+remoteJid, wac, msg)
		}

		time.Sleep(config.Get.Whatsapp.SendBDelay * time.Second)
	}

	message.Reply("Successfully send broadcasts messages", wac, msg)
	return nil
}
