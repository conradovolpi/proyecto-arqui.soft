package controllers

import (
	"backend/clients"
	"backend/dto"
	"backend/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	// Bind the request body to the SignUpRequest struct
	var body dto.SignUpRequest
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to bind JSON: %s", err.Error()),
		})
		return
	}
	// Call the Signup service
	err = service.UsuarioServiceInterfaceInstance.Signup(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to signup: %s", err.Error())})
		return
	}

	// Return the response
	result := dto.SignUpResponse{Message: "User created successfully"}
	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	var client *dto.UserLoginRequest
	// Bind the request body to the LoginRequest struct
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Login service
	response := service.UsuarioServiceInterfaceInstance.Login(*client)
	if response.Message == "Invalid email or password" || response.Message == "Error generating token" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": response.Message})
		return
	}

	userDAO, _ := clients.ObtainUserByEmail(client.Email) // Obtener el usuario para obtener el userId

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", response.Token, 60*60*24*30, "/", "localhost", false, true)
	c.SetCookie("Auth", response.Token, 60*60*24*30, "", "", false, true)
	c.SetCookie("userId", strconv.Itoa(int(userDAO.ID)), 60*60*24*30, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Cookie successfully generated"})
}
