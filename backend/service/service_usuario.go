package service

import (
	"backend/dao"

	"gorm.io/gorm"
)

type UsuarioService struct {
	db *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{db: db}
}

// CrearUsuario delega al DAO la creación del usuario
func (s *UsuarioService) CrearUsuario(usuario *dao.Usuario) error {
	return dao.CrearUsuario(s.db, usuario)
}

// ObtenerUsuarioPorID recupera un usuario por su ID
func (s *UsuarioService) ObtenerUsuarioPorID(id int) (*dao.Usuario, error) {
	return dao.ObtenerUsuarioPorID(s.db, id)
}

// EliminarUsuario elimina un usuario dado su ID
func (s *UsuarioService) EliminarUsuario(id int) error {
	return dao.EliminarUsuario(s.db, id)
}

// ObtenerUsuarioPorEmail (opcional) si querés validación o login
func (s *UsuarioService) ObtenerUsuarioPorEmail(email string) (*dao.Usuario, error) {
	return dao.ObtenerUsuarioPorEmail(s.db, email)
}
