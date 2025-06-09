package inscripcion

import (
	"backend/clients/inscripcion"
	"backend/dto"
	"backend/models"
	"backend/utils"
	"time"
)

type inscripcionService struct{}

var InscripcionService inscripcionServiceInterface

type inscripcionServiceInterface interface {
	Inscribir(dto.InscripcionCreateDTO) (dto.InscripcionResponseDTO, utils.ApiError)
	Cancelar(uint, uint) utils.ApiError
}

func init() {
	InscripcionService = &inscripcionService{}
}

func (s *inscripcionService) Inscribir(input dto.InscripcionCreateDTO) (dto.InscripcionResponseDTO, utils.ApiError) {
	insc := models.Inscripcion{
		UsuarioID:        input.UsuarioID,
		ActividadID:      input.ActividadID,
		FechaInscripcion: time.Now(),
	}

	err := inscripcion.InscripcionClient.Create(&insc)
	if err != nil {
		return dto.InscripcionResponseDTO{}, utils.NewInternalServerApiError("Error registrando inscripción", err)
	}

	return dto.InscripcionResponseDTO{
		UsuarioID:        insc.UsuarioID,
		ActividadID:      insc.ActividadID,
		FechaInscripcion: insc.FechaInscripcion,
	}, nil
}

func (s *inscripcionService) Cancelar(usuarioID uint, actividadID uint) utils.ApiError {
	err := inscripcion.InscripcionClient.Delete(usuarioID, actividadID)
	if err != nil {
		return utils.NewInternalServerApiError("Error cancelando inscripción", err)
	}
	return nil
}
