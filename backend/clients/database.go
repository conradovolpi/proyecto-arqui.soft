package clients

import (
	"backend/dao"
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() error {
	// Leer la configuración desde el entorno
	dsn := os.Getenv("DB")
	if dsn == "" {
		return fmt.Errorf("database connection string is empty")
	}

	// Conectar a la base de datos
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Registrar éxito
	log.Println("Successfully connected to the database")

	// Migraciones
	/*
		err = DB.AutoMigrate(&dao.Course{}, &dao.Category{}, &dao.User{}, &dao.CourseInscription{}, &dao.Comment{})
		if err != nil {
			return fmt.Errorf("failed to migrate database: %w", err)
		}
	*/
	return nil
}

// USUARIOS
func CreateUser(user *dao.Usuario) error {
	result := DB.Create(user)
	return result.Error
}

// update user
// ACTIVIDADES
func CreateActivity(actividad dao.Actividad) error {
	result := DB.Create(&actividad)
	if result.Error != nil {
		log.Println("Error al crear actividad:", result.Error)
	}
	return result.Error
}

func GetAllActivities() ([]dao.Actividad, error) {
	var actividad []dao.Actividad
	result := DB.Find(&actividad)
	if result.Error != nil {
		return nil, result.Error
	}
	return actividad, nil
}

func ObtainActivityByName(titulo string) (*dao.Actividad, error) {
	var actividad dao.Actividad
	result := DB.Where("titulo = ?", titulo).
		First(&actividad)
	if result.Error != nil {
		return nil, result.Error
	}
	return &actividad, nil
}

func UpdateActivityByID(id uint, actividad dao.Actividad) error {
	result := DB.Model(&dao.Actividad{}).Where("id = ?", id).Updates(actividad)
	if result.Error != nil {
		log.Println("Error actualizando la actividad:", result.Error)
	}
	return result.Error
}

func DeleteActividadByID(id uint) error {
	result := DB.Delete(&dao.Actividad{}, id)
	if result.Error != nil {
		log.Println("Error al borrar la actividad:", result.Error)
	}
	return result.Error
}

// a chequear
func SearchActivity(query string) ([]dao.Actividad, error) {
	var actividad []dao.Actividad
	err := DB.Preload("Categories").
		Where("courses.name LIKE ? OR categories.name LIKE ?", "%"+query+"%", "%"+query+"%").
		Joins("JOIN course_categories ON courses.id = course_categories.course_id").
		Joins("JOIN categories ON categories.id = course_categories.category_id").
		Group("courses.id").
		Find(&actividad).Error

	if err != nil {
		return nil, err
	}

	if len(actividad) == 0 {
		return nil, errors.New("no se encontro la actividad")
	}

	return actividad, nil
}

//INSCRIPCIONES
/*
func EnrollUser(inscription dao.CourseInscription) error {
	result := DB.Create(&inscription)
	if result.Error != nil {
		return errors.New("error enrolling user: " + result.Error.Error())
	}
	return nil
}*/
