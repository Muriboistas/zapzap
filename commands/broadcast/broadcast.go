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
	commands.New("broadcast", broadcast).SetHelp("for broadcast management").SetCooldown(1).OnlyRoot().Add()
}

func broadcast(wac *whatsapp.Conn, msg whatsapp.TextMessage) error {
	var res string
	msgFields := strings.Fields(msg.Text)
	operation := msgFields[0]
	switch operation {
	case "list":
		res = formatBroadcastList()
	case "deactivate":
		broadcastID := msgFields[1]
		res = mapFromTo(broadcastID, operation, bc.Active, bc.Deactivated)
	case "activate":
		broadcastID := msgFields[1]
		res = mapFromTo(broadcastID, operation, bc.Deactivated, bc.Active)
	default:
		res = "Invalid operation"
	}

	message.Reply(res, wac, msg)

	return nil
}

func formatBroadcastList() string {
	list := []string{"Broadcasts", "Activated:"}
	for k := range bc.Active {
		list = append(list, fmt.Sprintf("*ID:* %s", k))
	}
	list = append(list, "\nDeactivated:")
	for k := range bc.Deactivated {
		list = append(list, fmt.Sprintf("*ID:* %s", k))
	}

	return strings.Join(list, "\n")
}

func mapFromTo(broadcastID, operation string, from, to map[string]string) string {
	v, found := from[broadcastID]
	if !found {
		return fmt.Sprintf("That broadcast doesn't exists in %sd list!", operation)
	}
	to[broadcastID] = v
	delete(from, broadcastID)
	return formatBroadcastList()
}
