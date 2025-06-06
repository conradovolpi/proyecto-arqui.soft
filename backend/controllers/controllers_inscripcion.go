package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/dao"
	"backend/service"

	"github.com/gorilla/mux"
)

type InscripcionController struct {
	Service *service.InscripcionService
}

func NewInscripcionController(service *service.InscripcionService) *InscripcionController {
	return &InscripcionController{Service: service}
}

func (c *InscripcionController) CrearInscripcion(w http.ResponseWriter, r *http.Request) {
	var inscripcion dao.Inscripcion
	if err := json.NewDecoder(r.Body).Decode(&inscripcion); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := c.Service.CrearInscripcion(&inscripcion); err != nil {
		http.Error(w, "Error al crear inscripción: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inscripcion)
}

func (c *InscripcionController) EliminarInscripcion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioID, err := strconv.Atoi(vars["usuarioID"])
	if err != nil {
		http.Error(w, "UsuarioID inválido", http.StatusBadRequest)
		return
	}
	actividadID, err := strconv.Atoi(vars["actividadID"])
	if err != nil {
		http.Error(w, "ActividadID inválido", http.StatusBadRequest)
		return
	}

	if err := c.Service.EliminarInscripcion(uint(usuarioID), uint(actividadID)); err != nil {
		http.Error(w, "Error al eliminar inscripción: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *InscripcionController) ObtenerInscripcionesPorUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioID, err := strconv.Atoi(vars["usuarioID"])
	if err != nil {
		http.Error(w, "UsuarioID inválido", http.StatusBadRequest)
		return
	}

	inscripciones, err := c.Service.ObtenerInscripcionesPorUsuario(uint(usuarioID))
	if err != nil {
		http.Error(w, "Error al obtener inscripciones: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inscripciones)
}
