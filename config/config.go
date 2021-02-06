package config

import (
	"log"
	"path/filepath"
	"runtime"
	"time"

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
		SessionPath:     pathFromProjectRoot("session"),
		LongClientName:  "Muriboistas",
		ShortClientName: "Muriboistas",
		ClientVersion:   "1.0",
		RootNumber:      "",
	},
	Qrcode: Qrcode{
		FileName:    pathFromProjectRoot("session"),
		Quality:     "medium",
		Size:        256,
		GeneratePNG: true,
		PrintOnCLI:  true,
	},
	Database: Database{
		Path:           pathFromProjectRoot("data"),
		SaveBackup:     true,
		BackupPath:     pathFromProjectRoot("data/backups"),
		MigrationsPath: pathFromProjectRoot("data/migrations"),
	},
}

func loadConfig() Configuration {
	if err := configor.Load(&conf, pathFromProjectRoot("config/config.json")); err != nil {
		log.Println(err)
	}
	return conf
}

func pathFromProjectRoot(path string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return filepath.Join(basepath, "..", path)
}
