// controllers/controllers_inscripcion.go
package controllers

import (
	"net/http"

	"tu_modulo/dto"
	"tu_modulo/service"

	"github.com/gin-gonic/gin"
)

type InscripcionController struct {
	Service *service.InscripcionService
}

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

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Inscripción realizada con éxito"})
}
