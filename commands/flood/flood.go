package flood

import (
	"math/rand"
	"time"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New("flood", flood).SetHelp("flood some message").SetCooldown(8).Add()
}

func flood(wac *whatsapp.Conn, msg whatsapp.TextMessage) error {
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(5)
	message.Reply(msg.Text, wac, msg)
	for i := 0; i <= times+7; i++ {
		message.Send(msg.Text, wac, msg)
	}

	return nil
}
