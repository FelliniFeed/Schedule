package service

import "schedule/pkg/repository"

type SubjectService struct {
	repo repository.Subject
}

func NewSubjectService(repo repository.Subject) *SubjectService{
	return &SubjectService{repo:repo}
}