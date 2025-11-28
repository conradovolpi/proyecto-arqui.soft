package models

type Usuario struct {
	UsuarioID     uint          `gorm:"primaryKey" json:"usuario_id"`
	Nombre        string        `json:"nombre"`
	Email         string        `gorm:"uniqueIndex:idx_usuarios_email,length:255" json:"email"`
	Password      string        `gorm:"type:varchar(64)" json:"password"` // SHA256 produce 64 caracteres hex
	Rol           string        `json:"rol"`
	Inscripciones []Inscripcion `gorm:"foreignKey:UsuarioID" json:"inscripciones,omitempty"`
}
