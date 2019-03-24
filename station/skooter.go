package station

import (
	"github.com/globalsign/mgo/bson"
	"muskooters/services/mongo"
	"muskooters/services/assert"
	"github.com/globalsign/mgo"
	"errors"
)

const collectionName = "scooters"

var starterState = Ready

type Scooter struct {
	ID    bson.ObjectId `bson:"_id"`
	State State
}

func NewScooter() Scooter {
	id := bson.NewObjectId()
	scooter := Scooter{ID: id, State: starterState}

	c := mongo.GetDB().C(collectionName)
	err := c.Insert(scooter)
	assert.Nil(err)

	return scooter
}

