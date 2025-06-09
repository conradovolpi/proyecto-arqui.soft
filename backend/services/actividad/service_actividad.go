package actividad

import (
	"backend/clients/actividad"
	"backend/dto"
	"backend/models"
	"backend/utils"
)

type actividadService struct{}

var ActividadService actividadServiceInterface

type actividadServiceInterface interface {
	Create(dto.ActividadCreateDTO) (dto.ActividadResponseDTO, utils.ApiError)
	GetByID(uint) (dto.ActividadResponseDTO, utils.ApiError)
	GetAll() ([]dto.ActividadResponseDTO, utils.ApiError)
}

func init() {
	ActividadService = &actividadService{}
}

func (s *actividadService) Create(input dto.ActividadCreateDTO) (dto.ActividadResponseDTO, utils.ApiError) {
	act := models.Actividad{
		Titulo:        input.Titulo,
		Descripcion:   input.Descripcion,
		Instructor:    input.Instructor,
		Cupo:          input.Cupo,
		Categoria:     input.Categoria,
		HorarioInicio: input.HorarioInicio,
		HorarioFin:    input.HorarioFin,
	}

	err := actividad.ActividadClient.Create(&act)
	if err != nil {
		return dto.ActividadResponseDTO{}, utils.NewInternalServerApiError("Error creando actividad", err)
	}

	return dto.ActividadResponseDTO{
		ActividadID:   act.ActividadID,
		Titulo:        act.Titulo,
		Descripcion:   act.Descripcion,
		Instructor:    act.Instructor,
		Cupo:          act.Cupo,
		Categoria:     act.Categoria,
		HorarioInicio: act.HorarioInicio,
		HorarioFin:    act.HorarioFin,
	}, nil
}

func (s *actividadService) GetByID(id uint) (dto.ActividadResponseDTO, utils.ApiError) {
	act, err := actividad.ActividadClient.GetByID(id)
	if err != nil {
		return dto.ActividadResponseDTO{}, utils.NewNotFoundApiError("Actividad no encontrada")
	}

	return dto.ActividadResponseDTO{
		ActividadID:   act.ActividadID,
		Titulo:        act.Titulo,
		Descripcion:   act.Descripcion,
		Instructor:    act.Instructor,
		Cupo:          act.Cupo,
		Categoria:     act.Categoria,
		HorarioInicio: act.HorarioInicio,
		HorarioFin:    act.HorarioFin,
	}, nil
}

func (s *actividadService) GetAll() ([]dto.ActividadResponseDTO, utils.ApiError) {
	acts, err := actividad.ActividadClient.GetAll()
	if err != nil {
		return nil, utils.NewInternalServerApiError("Error listando actividades", err)
	}

	var dtos []dto.ActividadResponseDTO
	for _, act := range acts {
		dtos = append(dtos, dto.ActividadResponseDTO{
			ActividadID:   act.ActividadID,
			Titulo:        act.Titulo,
			Descripcion:   act.Descripcion,
			Instructor:    act.Instructor,
			Cupo:          act.Cupo,
			Categoria:     act.Categoria,
			HorarioInicio: act.HorarioInicio,
			HorarioFin:    act.HorarioFin,
		})
	}
	return dtos, nil
}
