package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/muriboistas/zapzap/data"
	"github.com/muriboistas/zapzap/infra/whats"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var migrate = flag.String("migrate", "", `start migrations if param is equal "up" or "down"`)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006/01/02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               (func([]string))(nil),
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  logrus.FieldMap(nil),
		CallerPrettyfier:          (func(*runtime.Frame) (string, string))(nil),
	})
}

func main() {
	log.Trace("Parsing flags")
	flag.Parse()

	if *migrate != "" {
		err := data.Migration(*migrate)
		if err != nil {
			log.Error(err)
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
