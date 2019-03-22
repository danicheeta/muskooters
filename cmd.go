package main

import (
	"github.com/pressly/goose"
	"muskooters/services/assert"
	"muskooters/services/config"
	"muskooters/services/initializer"
	"muskooters/services/mysql"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"muskooters/user"
)

const appName = "muskooters"

func main() {
	config.Init(appName)
	defer initializer.Initialize()()

	migup()
	err := user.Add("daniel", "123", user.Hunter)
	assert.Nil(err)

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
