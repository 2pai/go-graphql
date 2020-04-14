package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	SecretKey = []byte("secret")
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	payload := token.Claims.(jwt.MapClaims)
	payload["username"] = username
	payload["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if payload, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := payload["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
