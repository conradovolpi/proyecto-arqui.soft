// controllers/controllers_inscripcion.go
package controllers

import (
	"net/http"
	"strconv"

	"backend/dto"
	"backend/service"

	"github.com/gin-gonic/gin"
)

type InscripcionController struct {
	Service *service.InscripcionService
}

// POST /inscripciones
func (ctrl *InscripcionController) PostInscripcion(c *gin.Context) {
	var input dto.CrearInscripcionDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := ctrl.Service.CrearInscripcion(input.UsuarioID, input.ActividadID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Inscripción exitosa"})
}

// GET /inscripciones
func (ctrl *InscripcionController) GetInscripciones(c *gin.Context) {
	inscripciones, err := ctrl.Service.ObtenerTodas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener inscripciones"})
		return
	}
	c.JSON(http.StatusOK, inscripciones)
}

// GET /inscripciones/usuario/:id
func (ctrl *InscripcionController) GetPorUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	inscripciones, err := ctrl.Service.ObtenerPorUsuario(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener inscripciones del usuario"})
		return
	}

	c.JSON(http.StatusOK, inscripciones)
}

// DELETE /inscripciones
func (ctrl *InscripcionController) DeleteInscripcion(c *gin.Context) {
	var input dto.CrearInscripcionDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := ctrl.Service.EliminarInscripcion(input.UsuarioID, input.ActividadID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Inscripción eliminada correctamente"})
}
