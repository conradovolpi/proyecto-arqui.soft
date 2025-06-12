package router

import (
	actividadCtrl "backend/controllers/actividad"
	inscripcionCtrl "backend/controllers/inscripcion"
	usuarioCtrl "backend/controllers/usuario"
	middleware "backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	usuarioController *usuarioCtrl.UsuarioController,
	inscripcionController *inscripcionCtrl.InscripcionController,
	actividadController *actividadCtrl.ActividadController,
) *gin.Engine {
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Ruta de ping para verificar si el servidor está vivo y CORS funciona
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Grupo de rutas para usuarios
	usuarios := r.Group("/usuarios")
	{
		usuarios.POST("/", usuarioController.Create)
		usuarios.POST("/login", usuarioController.Login)
		usuarios.GET("/", usuarioController.GetAll)
		usuarios.GET("/:id", usuarioController.GetByID)
	}

	// Grupo de rutas para inscripciones
	inscripciones := r.Group("/inscripciones")
	{
		inscripciones.POST("/", inscripcionController.Inscribir)
		inscripciones.DELETE("/", inscripcionController.Cancelar)
		inscripciones.GET("/usuario/:usuario_id", inscripcionController.GetPorUsuario)
		inscripciones.GET("/actividad/:actividad_id", inscripcionController.GetPorActividad)
	}

	// Grupo de rutas para actividades
	actividades := r.Group("/actividades")
	{
		actividades.Use(middleware.AdminOnly())
		actividades.GET("/", actividadController.GetAll)
		actividades.GET("/:id", actividadController.GetByID)
		actividades.POST("/", actividadController.Create)
		actividades.PUT("/:id", actividadController.Update)
		actividades.DELETE("/:id", actividadController.Delete)
	}

	return r
}
