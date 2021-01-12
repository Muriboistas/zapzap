package main

import (
	"fmt"

	"github.com/muriboistas/zapzap/infra/whats"
)

func main() {
	err := fmt.Errorf("qr code scan timed out")
	for err != nil {
		_, err = whats.New()
		if err != nil {
			fmt.Println(err)
		}
	}
}
