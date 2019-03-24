package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"muskooters/services/assert"
)

const secret = "muskuters+"

// generate token including just a role
func GenToken(role string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": role,
	})

	token, err := t.SignedString([]byte(secret))
	assert.Nil(err)

	return token
}

// fetch users role from token
func GetRoleFromToken(s string) (interface{}, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["role"].(string), nil
	}

	return "", errors.New("invalid token")
}
