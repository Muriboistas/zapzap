package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // sqlite3 drivers

	"github.com/muriboistas/zapzap/config"
	"github.com/muriboistas/zapzap/infra/repository/sqlite3"
	"github.com/muriboistas/zapzap/usecase/user"
)

// Service store services
var Service = loadServices()
var DB *sql.DB

// Services data
type Services struct {
	User *user.Service
}

func loadServices() *Services {
	s := &Services{}
	var err error
	DB, err = sql.Open("sqlite3", fmt.Sprintf("%s/db.sqlite", config.Get.Database.Path))
	if err != nil {
		log.Fatal(err)
	}

	userRepo := sqlite3.NewUser(DB)
	s.User = user.NewService(userRepo)

	return s
}
