package service

import (
	"backend/models"

	"gorm.io/gorm"
)

type UsuarioService struct {
	db *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{db: db}
}

func (s *UsuarioService) CrearUsuario(usuario *models.Usuario) error {
	return s.db.Create(usuario).Error
}

func (s *UsuarioService) ObtenerUsuarioPorID(id int) (*models.Usuario, error) {
	var usuario models.Usuario
	err := s.db.First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (s *UsuarioService) EliminarUsuario(id int) error {
	return s.db.Delete(&models.Usuario{}, id).Error
}
