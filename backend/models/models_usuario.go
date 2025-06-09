package models

type Usuario struct {
	UsuarioID     uint          `gorm:"primaryKey" json:"usuario_id"`
	Nombre        string        `json:"nombre"`
	Email         string        `gorm:"uniqueIndex" json:"email"`
	Password      string        `json:"password"`
	Rol           string        `json:"rol"`
	Inscripciones []Inscripcion `gorm:"foreignKey:UsuarioID" json:"inscripciones,omitempty"`
}
