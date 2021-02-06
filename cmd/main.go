package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/muriboistas/zapzap/data"
	"github.com/muriboistas/zapzap/infra/whats"
)

var migrate = flag.String("migrate", "", `start migrations if param is equal "up" or "down"`)

func main() {
	flag.Parse()

	if *migrate != "" {
		err := data.Migration(*migrate)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

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
	err = data.DB.Close()
	if err != nil {
		log.Fatalf("error closing database: %v\n", err)
	}
}
