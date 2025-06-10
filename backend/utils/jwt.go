package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(usuarioID uint, secret string) string {
	claims := jwt.MapClaims{
		"user_id": usuarioID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 1 día
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		panic("No se pudo firmar el token: " + err.Error())
	}
	return signedToken
}
