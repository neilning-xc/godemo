package token

import (
	"errors"
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

func Sign(c Context) (string, error) {
	secret := viper.GetString("jwt_secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": c.Username,
		"email":    c.Email,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func Parse(tokenString, jwtSecret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return ctx, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Username = claims["username"].(string)
		ctx.Email = claims["email"].(string)
		return ctx, nil
	}

	return ctx, err
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return &Context{}, errors.New("No Authrization")
	}

	var tokenString string
	fmt.Sscanf(header, "Bearer %s", &tokenString)
	jwtSecret := viper.GetString("jwt_secret")

	return Parse(tokenString, jwtSecret)
}
