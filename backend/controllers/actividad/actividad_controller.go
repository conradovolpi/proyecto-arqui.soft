package controllers

import (
	"backend/dto"
	actividad "backend/services/actividad"
	"backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActividadController struct {
	service actividad.ActividadServiceInterface
}

func NewActividadController(service actividad.ActividadServiceInterface) *ActividadController {
	return &ActividadController{service: service}
}

func (ac *ActividadController) Create(c *gin.Context) {
	var dto dto.ActividadCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("Datos inválidos"))
		return
	}

	result, err := ac.service.CrearActividad(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerApiError("Error creando actividad", err))
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (ac *ActividadController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("ID inválido"))
		return
	}

	result, err := ac.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewNotFoundApiError("Actividad no encontrada"))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ac *ActividadController) GetAll(c *gin.Context) {
	log.Printf("Iniciando GetAll de actividades")
	result, err := ac.service.GetAll()
	if err != nil {
		log.Printf("Error obteniendo actividades: %v", err)
		if err == gorm.ErrRecordNotFound {
			log.Printf("No se encontraron actividades en la base de datos")
			c.JSON(http.StatusOK, []dto.ActividadResponseDTO{})
			return
		}
		// En caso de error, devolver array vacío en lugar de error para evitar problemas en el frontend
		log.Printf("Error al obtener actividades, devolviendo array vacío: %v", err)
		c.JSON(http.StatusOK, []dto.ActividadResponseDTO{})
		return
	}

	// Asegurar que siempre sea un array (nunca nil)
	if result == nil {
		log.Printf("Resultado es nil, devolviendo array vacío")
		result = []dto.ActividadResponseDTO{}
	}

	log.Printf("Actividades obtenidas exitosamente: %d actividades", len(result))
	c.JSON(http.StatusOK, result)
}

func (ac *ActividadController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("ID inválido"))
		return
	}

	var dto dto.ActividadCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("Datos inválidos"))
		return
	}

	if err := ac.service.Update(uint(id), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerApiError("Error actualizando actividad", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Actividad actualizada"})
}

func (ac *ActividadController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("ID inválido"))
		return
	}

	if err := ac.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerApiError("Error eliminando actividad", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Actividad eliminada"})
}
