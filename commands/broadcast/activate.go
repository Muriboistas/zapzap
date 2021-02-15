package broadcast

import (
	"fmt"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/muriboistas/zapzap/commands"
	bc "github.com/muriboistas/zapzap/infra/whats/broadcast"
	"github.com/muriboistas/zapzap/infra/whats/message"
	"github.com/muriboistas/zapzap/pkg/helper/mapx"
)

func init() {
	commands.New(
		"broadcast", broadcastActivate,
	).SetSubcommands(
		"activate",
	).SetArgs(
		"broadcastID",
	).SetHelp(
		"for broadcast management",
	).SetCooldown(1).OnlyRoot().Add()
}

func broadcastActivate(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	var res string
	broadcastID := args["broadcastID"]
	if broadcastID == "" {
		res = "you can't activate blank spaces"
	}
	err := mapx.FromTo(broadcastID, &bc.Deactivated, &bc.Active)
	if err != nil {
		res = err.Error()
	} else {
		res = fmt.Sprintf("Successfully activate %s", broadcastID)
	}

	message.Reply(res, wac, msg)
	return nil
}
