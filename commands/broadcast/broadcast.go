package broadcast

import (
	"fmt"
	"strings"

	"github.com/muriboistas/zapzap/commands"

	bc "github.com/muriboistas/zapzap/infra/whats/broadcast"
	"github.com/muriboistas/zapzap/infra/whats/message"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New(
		"broadcast", broadcast,
	).SetHelp(
		"for broadcast management",
	).SetCooldown(1).OnlyRoot().Add()
}

func broadcast(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	list := []string{"Broadcasts"}
	if len(bc.Active) > 0 {
		list = append(list, "\nActive:")
		for k := range bc.Active {
			list = append(list, fmt.Sprintf("*ID:* %s", k))
		}
	}
	if len(bc.Deactivated) > 0 {
		list = append(list, "\nDeactivated:")
		for k := range bc.Deactivated {
			list = append(list, fmt.Sprintf("*ID:* %s", k))
		}
	}

	res := strings.Join(list, "\n")

	message.Reply(res, wac, msg)

	return nil
}
