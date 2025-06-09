package actividad

import (
	"backend/clients"
	"backend/models"
)

type ActividadClient struct{}

func NewActividadClient() *ActividadClient {
	return &ActividadClient{}
}

func (c *ActividadClient) Create(act *models.Actividad) error {
	return clients.Db.Create(act).Error
}

func (c *ActividadClient) GetByID(id uint) (*models.Actividad, error) {
	var act models.Actividad
	err := clients.Db.First(&act, id).Error
	return &act, err
}

func (c *ActividadClient) GetAll() ([]models.Actividad, error) {
	var acts []models.Actividad
	err := clients.Db.Find(&acts).Error
	return acts, err
}

func (c *ActividadClient) Update(act *models.Actividad) error {
	return clients.Db.Save(act).Error
}

func (c *ActividadClient) Delete(id uint) error {
	return clients.Db.Delete(&models.Actividad{}, id).Error
}
