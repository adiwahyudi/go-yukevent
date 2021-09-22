package middlewares

import (
	"time"
	"yukevent/constants"

	"github.com/golang-jwt/jwt"
)

func CreateTokenJWT(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRECT_JWT))
}
