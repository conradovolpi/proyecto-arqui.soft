package controllers

import (
	"backend/dto"
	usuario "backend/services/usuario"
	"backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	service usuario.UsuarioService
}

func NewUsuarioController(service usuario.UsuarioService) *UsuarioController {
	return &UsuarioController{service: service}
}

func (uc *UsuarioController) Create(c *gin.Context) {
	var userDTO dto.UsuarioCreateDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("Datos de entrada inválidos"))
		return
	}

	result, apiErr := uc.service.Create(userDTO)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (uc *UsuarioController) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("Credenciales inválidas"))
		return
	}

	result, apiErr := uc.service.Login(loginDTO)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (uc *UsuarioController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequestApiError("ID inválido"))
		return
	}

	result, apiErr := uc.service.GetByID(uint(id))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (uc *UsuarioController) GetAll(c *gin.Context) {
	result, apiErr := uc.service.GetAll()
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
