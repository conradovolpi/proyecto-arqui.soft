package inscripciones

import (
	"backend/models"
	"backend/services/inscripciones"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// POST /inscripciones
func InscribirUsuario(c *gin.Context) {
	var inscripcion models.Inscripcion

	if err := c.BindJSON(&inscripcion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := inscripciones.InscribirUsuarioEnActividad(inscripcion.UsuarioID, inscripcion.ActividadID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Inscripción exitosa"})
}

// GET /usuarios/:id/actividades
func GetActividadesDeUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	actividades, err := inscripciones.GetActividadesDeUsuario(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, actividades)
}
