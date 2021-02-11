package flip

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New("flip", flip).SetHelp("Flip a coin").SetCooldown(5).Add()
}

func flip(wac *whatsapp.Conn, msg whatsapp.TextMessage) error {
	var res string
	x := rand.Intn(2)

	v := strings.Fields(msg.Text)
	if len(v) != 2 {
		return errors.New("Invalid args")
	}

	side, err := strconv.Atoi(v[0])
	if err != nil || (side < 0 || side > 1) {
		return errors.New("Invalid coin side, use just 0 or 1")
	}
	bet, err := strconv.Atoi(v[1])
	if err != nil || bet < 1 {
		return errors.New("Invalid bet, use just numbers highter than 0")
	}

	if x == side {
		res = fmt.Sprintf("You win %d coins 🤑", bet*2)
	} else {
		res = fmt.Sprintf("You lose %d coins 💸", bet)
	}

	message.Reply(res, wac, msg)
	return nil
}
