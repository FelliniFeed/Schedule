package service

import "schedule/pkg/repository"

type StudentInClassService struct {
	repo repository.StudentInClass
}

func NewStudenInClassService(repo repository.StudentInClass) *StudentInClassService {
	return &StudentInClassService{repo: repo}
}