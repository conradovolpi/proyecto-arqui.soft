package controllers

import (
	"backend/dto"
	inscripcion "backend/services/inscripcion"
	"backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InscripcionController struct {
	service inscripcion.InscripcionServiceInterface
}

func NewInscripcionController(service inscripcion.InscripcionServiceInterface) *InscripcionController {
	return &InscripcionController{service: service}
}

func (ic *InscripcionController) Inscribir(c *gin.Context) {
	// Obtener el ID del usuario del token JWT
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedApiError("Usuario no autenticado"))
		return
	}

	var dto dto.InscripcionCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("Datos inválidos"))
		return
	}

	// Asignar el ID del usuario del token
	dto.UsuarioID = userID.(uint)

	if apiErr := ic.service.Inscribir(dto); apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inscripción realizada con éxito"})
}

func (ic *InscripcionController) Cancelar(c *gin.Context) {
	// Obtener el ID del usuario del token JWT
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedApiError("Usuario no autenticado"))
		return
	}

	var dto dto.InscripcionCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("Datos inválidos"))
		return
	}

	// Verificar que el usuario del token coincida con el usuario de la petición
	if userID.(uint) != dto.UsuarioID {
		c.JSON(http.StatusForbidden, utils.NewForbiddenApiError("No tienes permiso para cancelar esta inscripción"))
		return
	}

	if apiErr := ic.service.Cancelar(dto); apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción cancelada con éxito"})
}

func (ic *InscripcionController) GetPorUsuario(c *gin.Context) {
	idStr := c.Param("usuario_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("ID inválido"))
		return
	}

	result, apiErr := ic.service.GetPorUsuario(uint(id))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ic *InscripcionController) GetPorActividad(c *gin.Context) {
	idStr := c.Param("actividad_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("ID inválido"))
		return
	}

	result, apiErr := ic.service.GetPorActividad(uint(id))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
