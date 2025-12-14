package services

import (
	"context"
	"rup_database/internal/models"
	"rup_database/internal/repository"
)

type ObjectService struct {
	repo repository.ObjectRepository
}

func NewObjectService(repo repository.ObjectRepository) *ObjectService {
	return &ObjectService{repo: repo}
}

func (s *ObjectService) Create(ctx context.Context, input models.CreateObjectDTO) (*models.Object, error) {
	return s.repo.Create(ctx, input)
}

func (s *ObjectService) GetByID(ctx context.Context, id string) (*models.Object, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ObjectService) List(ctx context.Context) ([]*models.Object, error) {
	return s.repo.List(ctx)
}

func (s *ObjectService) Update(ctx context.Context, input models.UpdateObjectDTO) (*models.Object, error) {
	return s.repo.Update(ctx, input)
}

func (s *ObjectService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
