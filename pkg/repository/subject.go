package repository

import "github.com/jmoiron/sqlx"

type SubjectPostgres struct {
	db *sqlx.DB
}

func NewSubjectPostgres(db *sqlx.DB) *SchedulePostgres {
	return &SchedulePostgres{db:*db}
}