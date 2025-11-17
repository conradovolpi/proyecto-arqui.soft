package main

import (
	"backend/clients"
	"backend/clients/actividad"
	"backend/clients/inscripcion"
	actividadCtrl "backend/controllers/actividad"
	inscripcionCtrl "backend/controllers/inscripcion"
	usuarioCtrl "backend/controllers/usuario"
	"backend/router"
	actividadSvc "backend/services/actividad"
	inscripcionSvc "backend/services/inscripcion"
	usuarioSvc "backend/services/usuario"
)

func main() {
	// Conexión a base de datos y migración de entidades
	clients.ConnectDatabase()
	clients.MigrateEntities()

	// Inicialización de servicios
	usuarioService := usuarioSvc.NewUsuarioService()
	actividadService := actividadSvc.NewActividadService(actividad.ActividadClient)
	inscripcionService := inscripcionSvc.NewInscripcionService(inscripcion.InscripcionClient)

	// Inicialización de controladores
	usuarioController := usuarioCtrl.NewUsuarioController(usuarioService)
	actividadController := actividadCtrl.NewActividadController(actividadService)
	inscripcionController := inscripcionCtrl.NewInscripcionController(inscripcionService)

	// Seteo de rutas
	r := router.SetupRouter(usuarioController, actividadController, inscripcionController)

	// Iniciar el servidor
	r.Run(":8080") // Servidor corriendo en localhost:8080
}
