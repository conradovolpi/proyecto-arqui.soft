package service

import (
	dto "backend/dto/usuario"
	"backend/model"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte("super-secret-key") // 🔐 Cambialo por un valor más seguro en .env

type Claims struct {
	UsuarioID uint
	Rol       string
	jwt.RegisteredClaims
}

// Hashear contraseña con SHA256
func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// Registrar nuevo usuario
func RegisterUser(db *gorm.DB, req dto.UsuarioRequest) error {
	var existing model.Usuario
	result := db.Where("email = ?", req.Email).First(&existing)
	if result.Error == nil {
		return errors.New("el usuario ya existe")
	}

	hashedPassword := hashPassword(req.Password)

	user := model.Usuario{
		Nombre:   req.Nombre,
		Email:    req.Email,
		Password: hashedPassword,
		Rol:      req.Rol,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// Login y generar JWT
func LoginUser(db *gorm.DB, email, password string) (string, error) {
	var user model.Usuario
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("email o contraseña incorrectos")
	}

	if user.Password != hashPassword(password) {
		return "", errors.New("email o contraseña incorrectos")
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UsuarioID: user.UsuarioID,
		Rol:       user.Rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "gimnasio-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// Validar JWT y extraer datos
func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}
	return claims, nil
}

func ActualizarUsuario(db *gorm.DB, id uint, update dto.UsuarioRequest) error {
	var user model.Usuario
	if err := db.First(&user, id).Error; err != nil {
		return errors.New("usuario no encontrado")
	}

	// Solo actualizamos los campos permitidos
	user.Nombre = update.Nombre
	user.Email = update.Email
	user.Rol = update.Rol

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func EliminarUsuario(db *gorm.DB, id uint) error {
	var user model.Usuario
	if err := db.First(&user, id).Error; err != nil {
		return errors.New("usuario no encontrado")
	}

	if err := db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// creo que no es necesaria esta
// Obtener todos los usuarios
func ObtenerUsuarios(db *gorm.DB) ([]model.Usuario, error) {
	var usuarios []model.Usuario
	if err := db.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}
