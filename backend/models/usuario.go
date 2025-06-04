package models

type Usuario struct {
	UsuarioID     uint          `gorm:"primaryKey;autoIncrement"`
	Nombre        string        `gorm:"size:100;not null"`
	Email         string        `gorm:"size:100;not null;unique"`
	Password      string        `gorm:"size:256;not null"`
	Rol           string        `gorm:"size:20;not null"`
	Inscripciones []Inscripcion `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
}
