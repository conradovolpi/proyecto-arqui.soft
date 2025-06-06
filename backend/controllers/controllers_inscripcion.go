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

// POST /inscripciones
func (c *InscripcionController) CrearInscripcion(w http.ResponseWriter, r *http.Request) {
	var inscripcion dao.Inscripcion
	if err := json.NewDecoder(r.Body).Decode(&inscripcion); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := c.Service.CrearInscripcion(&inscripcion); err != nil {
		http.Error(w, "Error al crear inscripción", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inscripcion)
}

// GET /inscripciones/usuario/{usuarioID}
func (c *InscripcionController) ObtenerPorUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["usuarioID"]
	usuarioID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	inscripciones, err := c.Service.ObtenerInscripcionesPorUsuario(uint(usuarioID))
	if err != nil {
		http.Error(w, "Error al obtener inscripciones", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inscripciones)
}

// GET /inscripciones/actividad/{actividadID}
func (c *InscripcionController) ObtenerPorActividad(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["actividadID"]
	actividadID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID de actividad inválido", http.StatusBadRequest)
		return
	}

	inscripciones, err := c.Service.ObtenerInscripcionesPorActividad(uint(actividadID))
	if err != nil {
		http.Error(w, "Error al obtener inscripciones", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inscripciones)
}

// DELETE /inscripciones/{usuarioID}/{actividadID}
func (c *InscripcionController) EliminarInscripcion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioID, err1 := strconv.ParseUint(vars["usuarioID"], 10, 32)
	actividadID, err2 := strconv.ParseUint(vars["actividadID"], 10, 32)

	if err1 != nil || err2 != nil {
		http.Error(w, "IDs inválidos", http.StatusBadRequest)
		return
	}

	if err := c.Service.EliminarInscripcion(uint(usuarioID), uint(actividadID)); err != nil {
		http.Error(w, "Error al eliminar inscripción", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
