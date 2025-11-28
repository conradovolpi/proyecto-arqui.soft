package main

import (
	"backend/clients"
	"fmt"
	"log"
	"os"
)

func main() {
	// Configurar variables de entorno por defecto si no están definidas
	// (basado en docker-compose.yml)
	if os.Getenv("DB_USER") == "" {
		os.Setenv("DB_USER", "appuser")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "apppass")
	}
	if os.Getenv("DB_HOST") == "" {
		os.Setenv("DB_HOST", "localhost")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "3307")
	}
	if os.Getenv("DB_NAME") == "" {
		os.Setenv("DB_NAME", "appdb")
	}

	fmt.Println("Conectando a la base de datos...")
	fmt.Printf("Host: %s:%s, Database: %s, User: %s\n", 
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), 
		os.Getenv("DB_NAME"), os.Getenv("DB_USER"))

	// Conectar a la base de datos
	clients.ConnectDatabase()

	// Limpiar la base de datos
	if err := clients.ClearDatabase(); err != nil {
		log.Fatalf("Error al limpiar la base de datos: %v", err)
	}

	fmt.Println("\n============================================================")
	fmt.Println("⚠️  IMPORTANTE: Después de limpiar la base de datos:")
	fmt.Println("   1. Limpia el localStorage del navegador")
	fmt.Println("      (F12 -> Application -> Local Storage -> Clear)")
	fmt.Println("   2. O simplemente cierra sesión y vuelve a iniciar sesión")
	fmt.Println("   3. Los tokens JWT antiguos ya no serán válidos")
	fmt.Println("============================================================")
	fmt.Println("\nProceso completado exitosamente")
}

