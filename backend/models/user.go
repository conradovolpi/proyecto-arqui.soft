package models

type Usuario struct {
	UsuarioID     uint          `gorm:"primaryKey;autoIncrement"`
	Nombre        string        `gorm:"size:100;not null"`
	Email         string        `gorm:"size:100;unique;not null"`
	Password      string        `gorm:"size:256;not null"` // Aquí se almacena el hash
	Rol           string        `gorm:"size:20;not null"`  // "socio" o "admin"
	Inscripciones []Inscripcion `gorm:"foreignKey:UsuarioID"`
}
 
gola