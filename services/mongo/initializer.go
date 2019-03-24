package mongo

import (
	"github.com/globalsign/mgo"
	"muskooters/services/assert"
	"muskooters/services/config"
	"muskooters/services/initializer"
)

type initMongo struct{}

var muskooter *mgo.Database

// Initialize the modules, its safe to call this as many time as you want.
func (in initMongo) Initialize() func() {
	mongoURL := config.MustString("MONGO_URL")
	mongoUser := config.MustString("MONGO_USER")
	mongoPass := config.MustString("MONGO_PASS")
	mongoDBName := config.MustString("MONGO_DB_NAME")

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{mongoURL},
		Username: mongoUser,
		Password: mongoPass,
		Source:   "admin",
		FailFast: true,
	})
	assert.Nil(err)

	err = session.Ping()
	assert.Nil(err)

	muskooter = session.DB(mongoDBName)

	return func() {
		session.Close()
	}
}

func GetDB() *mgo.Database {
	return muskooter
}

func init() {
	initializer.Register(initMongo{})
}
