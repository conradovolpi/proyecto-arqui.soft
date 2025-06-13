package router

import (
	actividadCtrl "backend/controllers/actividad"
	inscripcionCtrl "backend/controllers/inscripcion"
	usuarioCtrl "backend/controllers/usuario"
	middleware "backend/middleware"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	usuarioController *usuarioCtrl.UsuarioController,
	inscripcionController *inscripcionCtrl.InscripcionController,
	actividadController *actividadCtrl.ActividadController,
) *gin.Engine {
	// Inicializar Gin y desactivar la redirección de barras diagonales
	r := gin.New()
	r.RedirectTrailingSlash = false

	// Añadir middleware de logging y recuperación de Gin
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Configuración de CORS: Aplicar aquí, antes de definir las rutas.
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

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

	// Grupo de rutas para inscripciones (requiere autenticación)
	inscripciones := r.Group("/inscripciones")
	inscripciones.Use(middleware.AuthRequired())
	{
		inscripciones.POST("/", inscripcionController.Inscribir)
		inscripciones.GET("/usuario/:usuario_id", inscripcionController.GetPorUsuario)
		inscripciones.GET("/actividad/:actividad_id", inscripcionController.GetPorActividad)
		inscripciones.DELETE("/", inscripcionController.Cancelar)
	}

	// Grupo de rutas para actividades
	actividades := r.Group("/actividades")
	{
		// Rutas públicas
		actividades.GET("/", actividadController.GetAll)
		actividades.GET("/:id", actividadController.GetByID)

		// Rutas que requieren ser admin
		actividadesAdmin := actividades.Group("")
		actividadesAdmin.Use(middleware.AdminOnly())
		{
			actividadesAdmin.POST("/", actividadController.Create)
			actividadesAdmin.PUT("/:id", actividadController.Update)
			actividadesAdmin.DELETE("/:id", actividadController.Delete)
		}
	}

	// Imprimir todas las rutas registradas
	routes := r.Routes()
	log.Println("Rutas registradas:")
	for _, route := range routes {
		log.Printf("%s %s", route.Method, route.Path)
	}

	return r
}
