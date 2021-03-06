package station

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"muskooters/services/assert"
	"muskooters/services/mongo"
)

const collectionName = "scooters"

var starterState = Ready

type Scooter struct {
	ID    bson.ObjectId `bson:"_id"`
	State State
}

// creates new scooter only by admin role permission
func NewScooter() Scooter {
	id := bson.NewObjectId()
	scooter := Scooter{ID: id, State: starterState}

	c := mongo.GetDB().C(collectionName)
	err := c.Insert(scooter)
	assert.Nil(err)

	return scooter
}

// fetch scooter from database
func GetScooter(id string) (Scooter, error) {
	objID := bson.ObjectIdHex(id)
	c := mongo.GetDB().C(collectionName)

	var s Scooter
	err := c.FindId(objID).One(&s)
	if err == mgo.ErrNotFound {
		return Scooter{}, errors.New("scooter not found")
	}
	assert.Nil(err)

	return Scooter{objID, s.State}, nil
}

// update scooter's state status
func SetScooterState(id bson.ObjectId, s State) error {
	c := mongo.GetDB().C(collectionName)
	return c.Update(bson.M{"_id": id}, bson.M{"state": s})

}

func getScootersWithBountyState() []Scooter {
	scooters := []Scooter{}

	c := mongo.GetDB().C(collectionName)
	if err := c.Find(bson.M{"state": Bounty}).All(&scooters); err != nil {
		assert.Nil(err)
	}

	return scooters
}

// updates state status in database
func (scooter Scooter) CommitTransit() {
	SetScooterState(scooter.ID, scooter.State)
}