package usuario

import (
	"backend/clients"
	"backend/models"
	"log"
)

type UsuarioClientInterface interface {
	GetByID(id uint) (*models.Usuario, error)
	GetByEmail(email string) (*models.Usuario, error)
	GetAll() ([]models.Usuario, error)
	CreateUser(usuario models.Usuario) (models.Usuario, error)
	UpdateUser(usuario models.Usuario) error
	DeleteUser(id uint) error
}

type usuarioClient struct{}

var UsuarioClient UsuarioClientInterface = &usuarioClient{}

func (u *usuarioClient) GetByID(id uint) (*models.Usuario, error) {
	var user models.Usuario
	err := clients.Db.First(&user, id).Error
	return &user, err
}

func (u *usuarioClient) GetByEmail(email string) (*models.Usuario, error) {
	var user models.Usuario
	err := clients.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usuarioClient) GetAll() ([]models.Usuario, error) {
	var users []models.Usuario
	err := clients.Db.Find(&users).Error
	return users, err
}

func (u *usuarioClient) CreateUser(usuario models.Usuario) (models.Usuario, error) {
	err := clients.Db.Create(&usuario).Error
	if err != nil {
		log.Println("Error creando usuario:", err)
	}
	return usuario, err
}

func (u *usuarioClient) UpdateUser(usuario models.Usuario) error {
	return clients.Db.Save(&usuario).Error
}

func (u *usuarioClient) DeleteUser(id uint) error {
	return clients.Db.Delete(&models.Usuario{}, id).Error
}
