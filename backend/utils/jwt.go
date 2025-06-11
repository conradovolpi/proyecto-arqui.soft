package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenInvalido = errors.New("token inválido")
	ErrTokenExpirado = errors.New("token expirado")
)

func GenerateJWT(usuarioID uint, rol, secret string) string {
	claims := jwt.MapClaims{
		"user_id": usuarioID,
		"rol":     rol,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		panic("No se pudo firmar el token: " + err.Error())
	}
	return signedToken
}

func ValidateJWT(tokenString string, secret string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalido
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, "", ErrTokenExpirado
		}
		return 0, "", ErrTokenInvalido
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, idOk := claims["user_id"].(float64)
		rol, rolOk := claims["rol"].(string)
		if idOk && rolOk {
			return uint(userID), rol, nil
		}
	}

	return 0, "", ErrTokenInvalido
}
