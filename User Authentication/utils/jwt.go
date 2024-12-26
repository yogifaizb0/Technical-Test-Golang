package utils

import (
	"time"
	"userauthentication/models"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("dhskh7iwye37das2dsadsda49saewquhjk78dlkasjdgi")

func GenerateToken(user models.User) (string, error){
	expTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{
		"username": user.Username,
		"expire": expTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string)(*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JWT_KEY, nil
	})
	return token, err
}