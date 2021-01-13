package main

import (
	"github.com/muriboistas/zapzap/infra/whats"
)

func main() {
	// reload the connection if fail
	for {
		_, err := whats.New()
		if err == nil {
			break
		}
	}
}
