package clients

import (
	"backend/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func ConnectDatabase() {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSchema := os.Getenv("DB_NAME")
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	connection := fmt.Sprintf(dsn, dbUsername, dbPassword, dbHost, dbPort, dbSchema)

	var err error
	for i := 0; i < 10; i++ {
		Db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
		if err == nil {
			fmt.Println("Conexión establecida con la base de datos")
			break
		}
		fmt.Println("Error conectando a la DB. Reintentando en 5s...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf(" No se pudo conectar a la base de datos: %v", err)
	}
}

func MigrateEntities() {
	err := Db.AutoMigrate(&models.Usuario{}, &models.Actividad{}, &models.Inscripcion{})
	if err != nil {
		log.Fatalf("Error migrando entidades: %v", err)
	}
	fmt.Println("Migración de entidades completada")
}
