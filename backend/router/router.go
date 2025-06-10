package router

import (
	actividadCtrl "backend/controllers/actividad"
	inscripcionCtrl "backend/controllers/inscripcion"
	usuarioCtrl "backend/controllers/usuario"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	usuarioController *usuarioCtrl.UsuarioController,
	inscripcionController *inscripcionCtrl.InscripcionController,
	actividadController *actividadCtrl.ActividadController,
) *gin.Engine {
	r := gin.Default()

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
		actividades.POST("/", actividadController.Create)
		actividades.GET("/", actividadController.GetAll)
		actividades.GET("/:id", actividadController.GetByID)
		actividades.PUT("/:id", actividadController.Update)
		actividades.DELETE("/:id", actividadController.Delete)
	}

	return r
}
