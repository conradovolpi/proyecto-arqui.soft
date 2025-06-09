package inscripcion

import (
	"backend/clients"
	"backend/models"
)

type InscripcionClientInterface interface {
	Get(usuarioID, actividadID uint) (*models.Inscripcion, error)
	GetByUsuario(usuarioID uint) ([]models.Inscripcion, error)
	GetByActividad(actividadID uint) ([]models.Inscripcion, error)
	Create(insc models.Inscripcion) error
	Delete(usuarioID, actividadID uint) error
}

type inscripcionClient struct{}

var InscripcionClient InscripcionClientInterface = &inscripcionClient{}

func (i *inscripcionClient) Get(usuarioID, actividadID uint) (*models.Inscripcion, error) {
	var insc models.Inscripcion
	err := clients.Db.Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).
		First(&insc).Error
	return &insc, err
}

func (i *inscripcionClient) GetByUsuario(usuarioID uint) ([]models.Inscripcion, error) {
	var insc []models.Inscripcion
	err := clients.Db.Where("usuario_id = ?", usuarioID).Find(&insc).Error
	return insc, err
}

func (i *inscripcionClient) GetByActividad(actividadID uint) ([]models.Inscripcion, error) {
	var insc []models.Inscripcion
	err := clients.Db.Where("actividad_id = ?", actividadID).Find(&insc).Error
	return insc, err
}

func (i *inscripcionClient) Create(insc models.Inscripcion) error {
	return clients.Db.Create(&insc).Error
}

func (i *inscripcionClient) Delete(usuarioID, actividadID uint) error {
	return clients.Db.Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).
		Delete(&models.Inscripcion{}).Error
}
