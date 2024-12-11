package service

import "schedule/pkg/repository"

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService{
	return &StudentService{repo:repo}
}