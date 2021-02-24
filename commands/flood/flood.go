package flood

import (
	"errors"
	"math/rand"
	"time"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New(
		"flood", flood,
	).SetArgs(
		"text",
	).SetHelp(
		"flood some message",
	).SetCooldown(8).Add()
}

func flood(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	text := args["text"]
	if text == "" {
		return errors.New("You can not send blank messages")
	}
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(5)
	message.Reply(text, wac, msg)
	for i := 0; i <= times+7; i++ {
		message.Send(text, wac, msg)
	}

	return nil
}
