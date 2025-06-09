package usuario

import (
	"backend/clients"
	"backend/models"
)

type UsuarioClient struct{}

func NewUsuarioClient() *UsuarioClient {
	return &UsuarioClient{}
}

func (c *UsuarioClient) Create(usuario *models.Usuario) error {
	return clients.Db.Create(usuario).Error
}

func (c *UsuarioClient) GetByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	err := clients.Db.First(&usuario, id).Error
	return &usuario, err
}

func (c *UsuarioClient) GetAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := clients.Db.Find(&usuarios).Error
	return usuarios, err
}

func (c *UsuarioClient) Update(usuario *models.Usuario) error {
	return clients.Db.Save(usuario).Error
}

func (c *UsuarioClient) Delete(id uint) error {
	return clients.Db.Delete(&models.Usuario{}, id).Error
}
