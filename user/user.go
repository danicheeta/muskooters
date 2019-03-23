package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"muskooters/services/assert"
	"muskooters/services/mysql"
)

const userTable = "users"

func Add(username, pass string, role Role) error {
	q := fmt.Sprintf("insert into %s (username, passwd, role) values (?,?,?)", userTable)
	_, err := mysql.GetDBMap().Exec(q, username, hashString(pass), role)
	if err != nil {
		return err
	}

	return nil
}

func GetByName(username string) (User, error) {
	var user User
	q := fmt.Sprintf("select * from %s where username=?", userTable)
	err := mysql.GetDBMap().SelectOne(&user, q, username)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func hashString(s string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	assert.Nil(err)

	return hash
}
