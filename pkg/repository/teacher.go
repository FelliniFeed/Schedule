package repository

import "github.com/jmoiron/sqlx"

type TeacherPostgres struct {
	db *sqlx.DB
}

func NewTeacherPostgres(db *sqlx.DB) *TeacherPostgres {
	return &TeacherPostgres{db:db}
}