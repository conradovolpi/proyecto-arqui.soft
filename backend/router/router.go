package router

import (
	actividadCtrl "backend/controllers/actividad"
	inscripcionCtrl "backend/controllers/inscripcion"
	usuarioCtrl "backend/controllers/usuario"
	middleware "backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	usuarioController *usuarioCtrl.UsuarioController,
	actividadController *actividadCtrl.ActividadController,
	inscripcionController *inscripcionCtrl.InscripcionController,
) *gin.Engine {
	router := gin.Default()

	// 🔁 Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 🔽 Rutas de health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
		})
	})

	// Rutas de usuario
	router.POST("/usuarios/", usuarioController.Create)
	router.POST("/usuarios/login", usuarioController.Login)
	router.GET("/usuarios/", usuarioController.GetAll)
	router.GET("/usuarios/:id", usuarioController.GetByID)

	// Rutas de inscripciones (con autenticación)
	inscripciones := router.Group("/inscripciones")
	inscripciones.Use(middleware.AuthRequired())
	{
		inscripciones.POST("/", inscripcionController.Inscribir)
		inscripciones.GET("/usuario/:usuario_id", inscripcionController.GetPorUsuario)
		inscripciones.GET("/actividad/:actividad_id", inscripcionController.GetPorActividad)
		inscripciones.DELETE("/", inscripcionController.Cancelar)
	}

	// Rutas de actividades (públicas)
	router.GET("/actividades/", actividadController.GetAll)
	router.GET("/actividades/:id", actividadController.GetByID)

	// Rutas de actividades que requieren ser admin
	actividadesAdmin := router.Group("/actividades")
	actividadesAdmin.Use(middleware.AdminOnly())
	{
		actividadesAdmin.POST("/", actividadController.Create)
		actividadesAdmin.PUT("/:id", actividadController.Update)
		actividadesAdmin.DELETE("/:id", actividadController.Delete)
	}

	return router
}
