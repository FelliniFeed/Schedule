package service

import "schedule/pkg/repository"

type TeacherService struct {
	repo repository.Teacher
}

func NewTeacherService(repo repository.Teacher) *TeacherService{
	return &TeacherService{repo: repo}
}