package mysql

import (
	"database/sql"

	"gopkg.in/gorp.v2"
	"muskooters/services/assert"
	"muskooters/services/config"
	"muskooters/services/initializer"
)

type initMysql struct{}

var dbmap *gorp.DbMap

// Initialize the modules, its safe to call this as many time as you want.
func (in initMysql) Initialize() func() {
	sqlURL := config.MustString("MYSQL_URL")

	db, err := sql.Open("mysql", sqlURL)
	assert.Nil(err)

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	err = db.Ping()
	assert.Nil(err)

	return func() {
		dbmap.Db.Close()
	}
}

func GetDBMap() *gorp.DbMap {
	return dbmap
}

func init() {
	initializer.Register(initMysql{})
}
