package main

import (
	"backend/clients"
	"backend/clients/usuario"
	"backend/controllers"
	"backend/router"
)

func main() {
	// 1. Conexión y migración de la base de datos
	clients.ConnectDatabase()
	clients.MigrateEntities()

	// 2. Inicializar DAOs / Clients
	usuarioClient := usuario.NewUsuarioClient(clients.Db)

	// 3. Inicializar Controladores
	usuarioController := controllers.NewUsuarioController(usuarioClient)

	// 4. Inicializar Router
	r := router.SetupRouter(usuarioController)

	// 5. Lanzar servidor web
	if err := r.Run(":8080"); err != nil {
		panic("Error al iniciar el servidor: " + err.Error())
	}
}
