package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // sqlite3 migration driver
	_ "github.com/golang-migrate/migrate/v4/source/file"      // file manager

	"github.com/muriboistas/zapzap/config"
)

// Migration make the migrations
func Migration(mType string) error {
	mType = strings.ToLower(mType)
	if !(mType == "up" || mType == "down") {
		return errors.New("invalid migration type")
	}

	if config.Get.Database.SaveBackup {
		err := backupMigration()
		if err != nil {
			return err
		}
	}

	err := startMigrations(mType)
	if err != nil {
		return err
	}

	return nil
}

func startMigrations(migrationType string) error {
	migrationFile := fmt.Sprintf("file://data/migrations")
	connStr := fmt.Sprintf("sqlite3://data/db.sqlite")

	m, err := migrate.New(migrationFile, connStr)
	if err != nil {
		return err
	}

	log.Println("Starting migrations...")
	switch migrationType {
	case "down":
		if err := m.Down(); err != nil {
			return err
		}
		log.Println("Migrations downed successfully")
	case "up":
		if err := m.Up(); err != nil {
			return err
		}
		log.Println("Migrations upped successfully")
	default:
		return fmt.Errorf("error: invalid migration type %q", migrationType)
	}

	return nil
}

func backupMigration() error {
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/db.sqlite", config.Get.Database.Path))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/backup-%v.sqlite", config.Get.Database.BackupPath, time.Now().UnixNano()), content, 0644)
	if err != nil {
		return err
	}

	return nil
}
