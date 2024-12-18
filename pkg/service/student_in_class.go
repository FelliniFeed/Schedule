package service

import (
	"schedule/models"
	"schedule/pkg/repository"

	"github.com/google/uuid"
)

type StudentInClassService struct {
	repo repository.StudentInClass
}

func NewStudenInClassService(repo repository.StudentInClass) *StudentInClassService {
	return &StudentInClassService{repo: repo}
}

func(s *StudentInClassService) GetStudentInClass(studentId, classId string)([]models.StudentInClass, error) {
	return s.repo.GetStudentInClass(studentId, classId)
}

func(s *StudentInClassService) CreateStudentInClass(studentInClass models.StudentInClass)(uuid.UUID, error) {
	return s.repo.CreateStudentInClass(studentInClass)
}