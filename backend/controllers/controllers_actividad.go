package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/dao"
	"backend/service"

	"github.com/gorilla/mux"
)

type ActividadController struct {
	Service *service.ActividadService
}

func NewActividadController(service *service.ActividadService) *ActividadController {
	return &ActividadController{Service: service}
}

// POST /actividades
func (c *ActividadController) CrearActividad(w http.ResponseWriter, r *http.Request) {
	var actividad dao.Actividad
	if err := json.NewDecoder(r.Body).Decode(&actividad); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := c.Service.CrearActividad(&actividad); err != nil {
		http.Error(w, "Error al crear actividad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(actividad)
}

// GET /actividades/{id}
func (c *ActividadController) ObtenerActividad(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	actividad, err := c.Service.ObtenerActividadPorID(id)
	if err != nil {
		http.Error(w, "Actividad no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(actividad)
}

// DELETE /actividades/{id}
func (c *ActividadController) EliminarActividad(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := c.Service.EliminarActividad(id); err != nil {
		http.Error(w, "Error al eliminar actividad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /actividades
func (c *ActividadController) ListarActividades(w http.ResponseWriter, r *http.Request) {
	actividades, err := c.Service.ListarActividades()
	if err != nil {
		http.Error(w, "Error al obtener actividades", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(actividades)
}
