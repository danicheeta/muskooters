package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
	"muskooters/services/assert"
	"muskooters/services/config"
	"muskooters/services/initializer"
	"muskooters/services/mysql"
	"muskooters/station"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"muskooters/user/routes"
)

const appName = "muskooters"

var (
	_ routes.Route
	_ station.Route
)

func main() {
	config.Init(appName)
	defer initializer.Initialize()()

	migup()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP)
	<-c
}

func migup() {
	goose.SetDialect("mysql")
	migrationDir := config.MustString("MIGRATION_DIR")
	if err := goose.Up(mysql.GetDBMap().Db, migrationDir); err != goose.ErrNoNextVersion {
		assert.Nil(err)
	}
	logrus.Info("migration is up successfully")
}
