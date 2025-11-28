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
	// Ya no se carga ningún archivo .env acá

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

// ClearDatabase elimina todos los registros de las tablas y resetea los auto_increment
func ClearDatabase() error {
	fmt.Println("Iniciando limpieza de base de datos...")

	// Desactivar verificación de foreign keys temporalmente
	if err := Db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		log.Printf("Advertencia: No se pudo desactivar foreign key checks: %v", err)
	}

	// Eliminar en orden: primero inscripciones (por las foreign keys)
	// Usar Session con AllowGlobalUpdate para permitir DELETE sin WHERE
	if err := Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Inscripcion{}).Error; err != nil {
		log.Printf("Error eliminando inscripciones: %v", err)
		Db.Exec("SET FOREIGN_KEY_CHECKS = 1")
		return fmt.Errorf("error al limpiar inscripciones: %w", err)
	}
	fmt.Println("✓ Inscripciones eliminadas")

	// Luego actividades
	if err := Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Actividad{}).Error; err != nil {
		log.Printf("Error eliminando actividades: %v", err)
		Db.Exec("SET FOREIGN_KEY_CHECKS = 1")
		return fmt.Errorf("error al limpiar actividades: %w", err)
	}
	fmt.Println("✓ Actividades eliminadas")

	// Finalmente usuarios
	if err := Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Usuario{}).Error; err != nil {
		log.Printf("Error eliminando usuarios: %v", err)
		Db.Exec("SET FOREIGN_KEY_CHECKS = 1")
		return fmt.Errorf("error al limpiar usuarios: %w", err)
	}
	fmt.Println("✓ Usuarios eliminados")

	// Reactivar verificación de foreign keys
	if err := Db.Exec("SET FOREIGN_KEY_CHECKS = 1").Error; err != nil {
		log.Printf("Advertencia: No se pudo reactivar foreign key checks: %v", err)
	}

	fmt.Println("Base de datos limpiada exitosamente (auto_increment reseteado)")
	return nil
}