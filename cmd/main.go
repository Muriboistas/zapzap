package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/muriboistas/zapzap/infra/whats"
)

func main() {
	var err error
	// reload the connection if fail
	for {
		log.Println("Trying connection...")
		err := whats.New()
		if err == nil {
			break
		}
	}

	// Get signal
	log.Println("Server is running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	//Disconnect safe
	log.Println("Shutting down now...")
	err = whats.Disconnect()
	if err != nil {
		log.Fatalf("error disconnecting: %v\n", err)
	}
}
