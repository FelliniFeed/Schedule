package repository

import "github.com/jmoiron/sqlx"

type StudenPostgres struct {
	db *sqlx.DB
}

func NewStudenPostgres(db *sqlx.DB) *StudenPostgres{
	return &StudenPostgres{db:db}
}