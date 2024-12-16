package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": userName,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	})
	os.Setenv("jwtKey", "Yeswanth")
	tokenString, err := token.SignedString(os.Getenv("jwtKey"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
