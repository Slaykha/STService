package helpers

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const SecretKey = "secret"

func CreateUserToken(userId string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    userId,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return ""
	}
	return token
}

func CreateID() (id string) {
	id = uuid.New().String()

	id = strings.ReplaceAll(id, "-", "")

	id = id[0:8]

	return
}
