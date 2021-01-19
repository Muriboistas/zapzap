package flood

import (
	"math/rand"
	"time"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New("flood", flood).SetHelp("flood some message").SetCooldown(2).Add()
}

func flood(wac *whatsapp.Conn, msg whatsapp.TextMessage) {
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(5)
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for i := 0; i <= times+7; i++ {
			select {
			case <-ticker.C:
				message.Reply(msg.Text, wac, msg)
			case <-quit:
				ticker.Stop()
				return
			}
		}
		ticker.Stop()
	}()
}
