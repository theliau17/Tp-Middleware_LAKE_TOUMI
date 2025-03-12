package resource

import "middleware/example/internal/models"

type Repository interface {
	GetAll() ([]models.Resource, error)
	GetByID(id string) (*models.Resource, error)
	Create(r *models.Resource) error
	Update(r *models.Resource) error
	Delete(id string) error
}

type Service interface {
	GetResources() ([]models.Resource, error)
	GetResourceByID(id string) (*models.Resource, error)
	CreateResource(r *models.Resource) error
	UpdateResource(r *models.Resource) error
	DeleteResource(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetResources() ([]models.Resource, error) {
	return s.repo.GetAll()
}

func (s *service) GetResourceByID(id string) (*models.Resource, error) {
	return s.repo.GetByID(id)
}

func (s *service) CreateResource(r *models.Resource) error {
	return s.repo.Create(r)
}

func (s *service) UpdateResource(r *models.Resource) error {
	return s.repo.Update(r)
}

func (s *service) DeleteResource(id string) error {
	return s.repo.Delete(id)
}
