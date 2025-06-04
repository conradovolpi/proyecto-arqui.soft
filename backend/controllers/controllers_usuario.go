package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/models"
	"backend/service"

	"github.com/gorilla/mux"
)

type UsuarioController struct {
	Service *service.UsuarioService
}

func NewUsuarioController(service *service.UsuarioService) *UsuarioController {
	return &UsuarioController{Service: service}
}

// POST /usuarios
func (c *UsuarioController) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := c.Service.CrearUsuario(&usuario); err != nil {
		http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

// GET /usuarios/{id}
func (c *UsuarioController) ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	usuario, err := c.Service.ObtenerUsuarioPorID(id)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(usuario)
}

// DELETE /usuarios/{id}
func (c *UsuarioController) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := c.Service.EliminarUsuario(id); err != nil {
		http.Error(w, "Error al eliminar el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
