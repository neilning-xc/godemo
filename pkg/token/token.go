package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Context struct {
	Username string
	Email    string
}

func Sign(c Context) (string, error) {
	secret := viper.GetString("jwt_secret")
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": c.Username,
		"Email":    c.Email,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
