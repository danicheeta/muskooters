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

func GetScooterState(id string) (State, error) {
	objID := bson.ObjectIdHex(id)
	c := mongo.GetDB().C(collectionName)

	var s Scooter
	err := c.FindId(objID).One(&s)
	if err == mgo.ErrNotFound {
		return 0, errors.New("scooter not found")
	}
	assert.Nil(err)

	return s.State, nil
}

//func SetScooterState(id string, s State) {
//	objID := bson.ObjectIdHex(id)
//	c := mongo.GetDB().C(collectionName)
//	err := c.Update(bson.M{"_id": objID}, bson.M{"state": s})
//	assert.Nil(err)
//}