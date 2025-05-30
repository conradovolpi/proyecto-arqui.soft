package controllers

import (
	"backend/models"
	service "backend/services/actividades"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Crear actividad (solo admin)
func CrearActividad(c *gin.Context) {
	var actividad models.Actividad
	if err := c.BindJSON(&actividad); err != nil {
		log.Error("Error al parsear actividad: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creada, err := service.CrearActividad(actividad)
	if err != nil {
		log.Error("Error al crear actividad: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, creada)
}

// Obtener todas las actividades
func GetAllActividades(c *gin.Context) {
	actividades, err := service.GetAllActividades()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actividades)
}

// Obtener actividades por nombre (search)
func GetActividadesByNombre(c *gin.Context) {
	nombre := c.Query("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "falta el parámetro 'nombre'"})
		return
	}

	actividades, err := service.GetActividadesByNombre(nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actividades)
}

// Obtener actividad por ID
func GetActividadByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	actividad, err := service.GetActividadByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "actividad no encontrada"})
		return
	}
	c.JSON(http.StatusOK, actividad)
}

// Actualizar actividad
func UpdateActividad(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var data models.Actividad
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := service.UpdateActividad(uint(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Eliminar actividad
func DeleteActividad(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := service.DeleteActividad(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "actividad eliminada correctamente"})
}
