package middleware

import (
	"backend/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permite las solicitudes OPTIONS (preflight) pasar sin autenticación
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, rol, err := utils.ValidateJWT(token, os.Getenv("JWT_SECRET_KEY"))
		if err != nil || rol != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Acceso denegado"})
			return
		}

		// Podés guardar los datos del usuario en el contexto
		c.Set("userID", userID)
		c.Set("rol", rol)
		c.Next()
	}
}
