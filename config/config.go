package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/configor"
	_ "github.com/mattn/go-sqlite3" // sqlite3 drivers
)

// Get configs
var Get = loadConfig()

// Command config
type Command struct {
	Prefix string
}

// Whatsapp config
type Whatsapp struct {
	TimeOutDuration time.Duration
	SessionPath     string
	LongClientName  string
	ShortClientName string
	ClientVersion   string
	RootNumber      string
	SendBDelay      time.Duration
}

// Qrcode config
type Qrcode struct {
	FileName    string
	Quality     string
	Size        uint
	GeneratePNG bool
	PrintOnCLI  bool
}

// Database config
type Database struct {
	Path           string
	SaveBackup     bool
	BackupPath     string
	MigrationsPath string
}

// Configuration data
type Configuration struct {
	Command  Command
	Whatsapp Whatsapp
	Qrcode   Qrcode
	Database Database
}

var conf = Configuration{
	Command: Command{
		Prefix: ".",
	},
	Whatsapp: Whatsapp{
		TimeOutDuration: 5,
		SessionPath:     "session",
		LongClientName:  "Muriboistas",
		ShortClientName: "Muriboistas",
		ClientVersion:   "1.0",
		RootNumber:      "",
		SendBDelay:      180,
	},
	Qrcode: Qrcode{
		FileName:    "session",
		Quality:     "medium",
		Size:        256,
		GeneratePNG: true,
		PrintOnCLI:  false,
	},
	Database: Database{
		Path:           "data",
		SaveBackup:     true,
		BackupPath:     "data/backups",
		MigrationsPath: "data/migrations",
	},
}

func loadConfig() Configuration {
	configFile := "config/config.json"

	os.TempDir()
	if err := configor.Load(&conf, configFile); err != nil {
		log.Println(err)
	}
	return conf
}
