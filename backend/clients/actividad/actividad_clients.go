package actividad

import (
	"backend/clients"
	"backend/models"
)

type ActividadClientInterface interface {
	GetByID(id uint) (*models.Actividad, error)
	GetAll() ([]models.Actividad, error)
	Create(act models.Actividad) (models.Actividad, error)
	Update(act models.Actividad) error
	Delete(id uint) error
}

type actividadClient struct{}

var ActividadClient ActividadClientInterface = &actividadClient{}

func (a *actividadClient) GetByID(id uint) (*models.Actividad, error) {
	var act models.Actividad
	err := clients.Db.First(&act, id).Error
	return &act, err
}

func (a *actividadClient) GetAll() ([]models.Actividad, error) {
	var acts []models.Actividad
	err := clients.Db.Find(&acts).Error
	return acts, err
}

func (a *actividadClient) Create(act models.Actividad) (models.Actividad, error) {
	err := clients.Db.Create(&act).Error
	return act, err
}

func (a *actividadClient) Update(act models.Actividad) error {
	return clients.Db.Save(&act).Error
}

func (a *actividadClient) Delete(id uint) error {
	return clients.Db.Delete(&models.Actividad{}, id).Error
}
