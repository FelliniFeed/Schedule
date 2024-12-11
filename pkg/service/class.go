package service

import (
	"schedule/models"
	"schedule/pkg/repository"

	"github.com/google/uuid"
)

type ClassService struct {
	repo repository.Class
}

func NewClassService(repo repository.Class) *ClassService {
	return &ClassService{repo: repo}
}

func(s *ClassService) GetClasses()([]models.Class, error) {
	return s.repo.GetClasses()
}

func(s *ClassService) CreateClass(class models.Class) (uuid.UUID, error) {
	return s.repo.CreateClass(class)
}
