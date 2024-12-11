package service

import "schedule/pkg/repository"

type TimepairService struct {
	repo repository.Timepair
}

func NewTimepairService (repo repository.Timepair) *TimepairService{
	return &TimepairService{repo: repo}
}