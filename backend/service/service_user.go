package service

import (
     "gorm.io/gorm"
    "proyecto2025/backend/models"
	"backend/dto"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	
)

// crear usuario
func CreateUsuario(db *gorm.DB, usuario *models.Usuario) error {
    return db.Create(usuario).Error
}

// editar usuario
func UpdateUsuario(db *gorm.DB, id uint, nuevosDatos *models.Usuario) error {
    var usuario models.Usuario
    if err := db.First(&usuario, id).Error; err != nil {
        return err
    }

    // Actualiza los campos necesarios
    usuario.Nombre = nuevosDatos.Nombre
    usuario.Email = nuevosDatos.Email
    usuario.Password = nuevosDatos.Password // Asegurate de que ya venga hasheado
    usuario.Rol = nuevosDatos.Rol

    return db.Save(&usuario).Error
}

//borrar usuairo
func DeleteUsuario(db *gorm.DB, id uint) error {
    return db.Delete(&models.Usuario{}, id).Error
}

//login

var jwtKey = []byte("clave_secreta_segura") // podés leerla de una variable de entorno

type Claims struct {
	UsuarioID uint   `json:"usuario_id"`
	Email     string `json:"email"`
	Rol       string `json:"rol"`
	jwt.RegisteredClaims
}

func LoginUsuario(db *gorm.DB, loginDto dto.LoginDto) (string, error) {
	var user models.Usuario
	if err := db.Where("email = ?", loginDto.Email).First(&user).Error; err != nil {
		return "", errors.New("usuario no encontrado")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	claims := &Claims{
		UsuarioID: user.UsuarioID,
		Email:     user.Email,
		Rol:       user.Rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}




