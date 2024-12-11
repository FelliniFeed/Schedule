package repository

import "github.com/jmoiron/sqlx"

type SchedulePostgres struct {
	db sqlx.DB
}

func NewSchedulePostgres(db *sqlx.DB) *SchedulePostgres {
	return &SchedulePostgres{db:*db}
}