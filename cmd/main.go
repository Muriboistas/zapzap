package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rhymen/go-whatsapp"
	"github.com/muriboistas/zapzap/infra/whats"
)

func main() {
	var wac *whatsapp.Conn
	var err error
	// reload the connection if fail
	for {
		log.Println("Trying connection...")
		wac, err = whats.New()
		if err == nil {
			break
		}
	}

	log.Println("Server is running...")

	// Get signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	//Disconnect safe
	log.Println("Shutting down now.")
	_, err = wac.Disconnect()
	if err != nil {
		log.Fatalf("error disconnecting: %v\n", err)
	}
}
