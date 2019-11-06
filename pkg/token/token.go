package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Context struct {
	Username string
	Email    string
}

func Sign(cxt *gin.Context, c Context) (tokenString string, err error) {
	secret := viper.GetString("jwt_secret")
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": c.Username,
		"email":    c.Email,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err = token.SignedString([]byte(secret))
	fmt.Println(tokenString, err)
	return
}
