package service

import (
	"schedule/models"
	"schedule/pkg/repository"

	"github.com/google/uuid"
)

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService{
	return &StudentService{repo:repo}
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.repo.GetStudents()
}

func (s *StudentService) CreateStudent(student models.Student) (uuid.UUID, error) {
	return s.repo.CreateStudent(student)
}