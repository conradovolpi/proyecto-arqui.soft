package inscripcion

import (
	"backend/clients"
	"backend/models"
)

type InscripcionClient struct{}

func NewInscripcionClient() *InscripcionClient {
	return &InscripcionClient{}
}

func (c *InscripcionClient) Create(insc *models.Inscripcion) error {
	return clients.Db.Create(insc).Error
}

func (c *InscripcionClient) Get(usuarioID uint, actividadID uint) (*models.Inscripcion, error) {
	var insc models.Inscripcion
	err := clients.Db.Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).
		First(&insc).Error
	return &insc, err
}

func (c *InscripcionClient) Delete(usuarioID, actividadID uint) error {
	return clients.Db.Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).
		Delete(&models.Inscripcion{}).Error
}

func (c *InscripcionClient) GetByUsuario(usuarioID uint) ([]models.Inscripcion, error) {
	var inscripciones []models.Inscripcion
	err := clients.Db.Where("usuario_id = ?", usuarioID).Find(&inscripciones).Error
	return inscripciones, err
}

func (c *InscripcionClient) GetByActividad(actividadID uint) ([]models.Inscripcion, error) {
	var inscripciones []models.Inscripcion
	err := clients.Db.Where("actividad_id = ?", actividadID).Find(&inscripciones).Error
	return inscripciones, err
}
