package clients

import (
	"gorm.io/gorm"
	"backend/models"
	"os"
	"fmt"
	"log"
	"time"
	"gorm.io/driver/mysql"

)

var {
	db *gorm.DB
}

func ConnectDB() { 
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}


func MigrateEntities() {
	db.AutoMigrate(&models.Usuario{}, &models.Actividad{}, &models.Inscripcion{})
	if err != nil {
		fmt.Println("Error al migrar las entidades: ", err)
		panic(err)
	}
	fmt.Println("Migraciones aplicadas correctamente")
}





