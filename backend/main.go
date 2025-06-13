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
	"log"

	// "github.com/gin-contrib/cors" // Eliminado: CORS se configura ahora en router.go
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

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
	// Obtener el motor de Gin
	r := router.SetupRouter(usuarioController, actividadController, inscripcionController)

	// La configuración de CORS se maneja ahora dentro de router.go, no aquí.
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:5173"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	// config.AllowCredentials = true
	// r.Use(cors.New(config))

	// Iniciar el servidor
	r.Run(":8080") // Servidor corriendo en localhost:8080
}
