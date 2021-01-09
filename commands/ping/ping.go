package ping

import (
	"fmt"

	"github.com/muriboistas/zapzap/commands"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New("ping", ping).SetHelp("check the command").SetCooldown(1).OnlyRoot().Add()
}

func ping(*whatsapp.Conn, whatsapp.TextMessage) {
	fmt.Println("Pong")
}
