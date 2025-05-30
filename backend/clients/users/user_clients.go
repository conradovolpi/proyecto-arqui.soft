package users

import (
	"backend/models"
)

type UserClient struct{}

type UserClientInterface interface {
	CreateUser(user models.Usuario) error
	GetUserByEmail(email string) (models.Usuario, error)
	GetUserByID(id uint) (models.Usuario, error)
	UpdateUser(user models.Usuario) error
	DeleteUser(id uint) error
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &UserClient{}
}

func (c *UserClient) GetUserByID(user models.Usuario) error {

}
