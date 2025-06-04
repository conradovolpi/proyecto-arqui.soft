package main

import (
	"backend/clients"
	"backend/models"
)

func main() {
	clients.ConnectDB()
	clients.MigrateEntities()
	users.UserClient.CreateUser(models.Usuario{
		Nombre: "Juan Perez",
	}
}