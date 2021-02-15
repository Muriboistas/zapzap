package flip

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/message"

	"github.com/Rhymen/go-whatsapp"
)

func init() {
	commands.New(
		"flip", flip,
	).SetArgs(
		"side", "bet",
	).SetHelp(
		"Flip a coin",
	).SetCooldown(5).Add()
}

func flip(wac *whatsapp.Conn, msg whatsapp.TextMessage, args map[string]string) error {
	var res string
	sideStr, betStr := args["side"], args["bet"]
	if sideStr == "" || betStr == "" {
		return errors.New("Invalid arguments")
	}

	x := rand.Intn(2)

	side, err := strconv.Atoi(sideStr)
	if err != nil || (side < 0 || side > 1) {
		return errors.New("Invalid coin side, use just 0 or 1")
	}
	bet, err := strconv.Atoi(betStr)
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
