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
		"broadcast", broadcastDeactivateHelp,
	).SetSubcommands(
		"deactivate",
	).SetArgs(
		"broadcastID",
	).SetHelp(
		"for broadcast management",
	).SetCooldown(1).OnlyRoot().Add()
}

func broadcastDeactivateHelp(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	var res string
	broadcastID := args["broadcastID"]
	if broadcastID == "" {
		res = "you can't deactivate blank spaces"
	}
	err := mapx.FromTo(broadcastID, &bc.Active, &bc.Deactivated)
	if err != nil {
		res = err.Error()
	} else {
		res = fmt.Sprintf("Successfully deactivate %s", broadcastID)
	}

	message.Reply(res, wac, msg)
	return nil
}
