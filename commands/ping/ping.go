package ping

import (
	"fmt"

	"github.com/muriboistas/zapzap/commands"

	"github.com/muriboistas/zapzap/config"
	"github.com/muriboistas/zapzap/infra/whats/message"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New("ping", ping).SetHelp("check the command").SetCooldown(1). /*.OnlyRoot()*/ Add()
}

func ping(wac *whatsapp.Conn, msg whatsapp.TextMessage) {
	fmt.Println(config.Get.Whatsapp.SourceNumber)
	message.Reply("Pong", wac, msg)
}
