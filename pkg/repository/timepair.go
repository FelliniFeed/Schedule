package repository

import "github.com/jmoiron/sqlx"

type TimepairPostgres struct {
	db *sqlx.DB
}

func  NewTimepairPostgres(db *sqlx.DB) *TimepairPostgres {
	return &TimepairPostgres{db:db}
}