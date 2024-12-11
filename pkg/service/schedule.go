package service

import "schedule/pkg/repository"

type ScheduleService struct {
	repo repository.Schedule
}

func NewScheduleService(repo repository.Schedule) *ScheduleService{
	return &ScheduleService{repo: repo}
}

