package middleware

import (
	"backend/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthRequired es un middleware que verifica que el usuario esté autenticado
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permite las solicitudes OPTIONS (preflight) pasar sin autenticación
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		log.Printf("Verificando autenticación para la ruta: %s", c.Request.URL.Path)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Printf("No se encontró el header de autorización")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token requerido",
				"status": http.StatusUnauthorized,
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			log.Printf("Token vacío después de remover 'Bearer '")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Formato de token inválido",
				"status": http.StatusUnauthorized,
			})
			return
		}

		userID, rol, err := utils.ValidateJWT(token, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			log.Printf("Error validando token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token inválido o expirado",
				"status": http.StatusUnauthorized,
			})
			return
		}

		// Guardar los datos del usuario en el contexto
		c.Set("userID", userID)
		c.Set("rol", rol)
		log.Printf("Autenticación exitosa para usuario ID: %d, rol: %s", userID, rol)
		c.Next()
	}
}

// AdminOnly es un middleware que verifica que el usuario sea administrador
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permite las solicitudes OPTIONS (preflight) pasar sin autenticación
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		log.Printf("Verificando rol de administrador para la ruta: %s", c.Request.URL.Path)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Printf("No se encontró el header de autorización")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token requerido",
				"status": http.StatusUnauthorized,
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			log.Printf("Token vacío después de remover 'Bearer '")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Formato de token inválido",
				"status": http.StatusUnauthorized,
			})
			return
		}

		userID, rol, err := utils.ValidateJWT(token, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			log.Printf("Error validando token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token inválido o expirado",
				"status": http.StatusUnauthorized,
			})
			return
		}

		if rol != "admin" {
			log.Printf("Usuario ID %d con rol %s intentó acceder a ruta de administrador", userID, rol)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":  "Acceso denegado: se requiere rol de administrador",
				"status": http.StatusForbidden,
			})
			return
		}

		// Guardar los datos del usuario en el contexto
		c.Set("userID", userID)
		c.Set("rol", rol)
		log.Printf("Acceso de administrador exitoso para usuario ID: %d", userID)
		c.Next()
	}
}
