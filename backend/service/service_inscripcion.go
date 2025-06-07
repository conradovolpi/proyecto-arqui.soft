package service

/*
import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/model"
	"backend/repository"
)

type InscripcionService struct {
	repo *repository.InscripcionRepository
}

func NewInscripcionService(repo *repository.InscripcionRepository) *InscripcionService {
	return &InscripcionService{
		repo: repo,
	}
}

func (s *InscripcionService) Create(ctx context.Context, inscripcion *model.Inscripcion) error {
	inscripcion.ID = uuid.New()
	inscripcion.CreatedAt = time.Now()
	inscripcion.UpdatedAt = time.Now()

	return s.repo.Create(ctx, inscripcion)
}

func (s *InscripcionService) GetByID(ctx context.Context, id uuid.UUID) (*model.Inscripcion, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *InscripcionService) Update(ctx context.Context, inscripcion *model.Inscripcion) error {
	inscripcion.UpdatedAt = time.Now()
	return s.repo.Update(ctx, inscripcion)
}

func (s *InscripcionService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *InscripcionService) List(ctx context.Context) ([]*model.Inscripcion, error) {
	return s.repo.List(ctx)
}



*/
