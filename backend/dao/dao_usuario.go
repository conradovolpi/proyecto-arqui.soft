package dao

import (
	"gorm.io/gorm"
)

// Usuario representa tanto el modelo de base de datos como el punto de acceso a datos.
type Usuario struct {
	UsuarioID uint   `gorm:"primaryKey;autoIncrement"`
	Nombre    string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;not null;unique"`
	Password  string `gorm:"size:256;not null"`
	Rol       string `gorm:"size:20;not null"`
}

// ===============================
// Funciones DAO
// ===============================

// CrearUsuario guarda un nuevo usuario en la base
func CrearUsuario(db *gorm.DB, usuario *Usuario) error {
	return db.Create(usuario).Error
}

// ObtenerUsuarioPorID busca un usuario por ID
func ObtenerUsuarioPorID(db *gorm.DB, id int) (*Usuario, error) {
	var usuario Usuario
	err := db.First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// EliminarUsuario elimina un usuario por ID
func EliminarUsuario(db *gorm.DB, id int) error {
	return db.Delete(&Usuario{}, id).Error
}

// ObtenerUsuarioPorEmail busca un usuario por email (útil para login o validación)
func ObtenerUsuarioPorEmail(db *gorm.DB, email string) (*Usuario, error) {
	var usuario Usuario
	err := db.Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
