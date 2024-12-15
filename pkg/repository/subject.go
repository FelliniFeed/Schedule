package repository

import "github.com/jmoiron/sqlx"

type SubjectPostgres struct {
	db *sqlx.DB
}

func NewSubjectPostgres(db *sqlx.DB) *SubjectPostgres {
	return &SubjectPostgres{db: db}
}