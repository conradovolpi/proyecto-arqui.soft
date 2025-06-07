package dao

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre   string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"size:256;not null"`
	Rol      string `gorm:"size:20;not null"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
