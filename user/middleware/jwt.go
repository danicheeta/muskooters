package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"muskooters/services/assert"
	"muskooters/user"
	"fmt"
	"github.com/pkg/errors"
)

const secret = "muskuters+"

func GenToken(role string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": role,
	})

	token, err := t.SignedString([]byte(secret))
	assert.Nil(err)

	return token
}

func getRoleFromToken(s string) (user.Role, error) {
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
		return user.Role(claims["role"].(string)), nil
	}

	return "", errors.New("invalid token")
}