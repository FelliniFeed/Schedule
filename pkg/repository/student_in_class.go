package repository

import (
	"github.com/jmoiron/sqlx"
)

type StudenInClassPostgres struct {
	db *sqlx.DB
}

func NewStudenInClassPostgres(db *sqlx.DB) *StudenInClassPostgres {
	return &StudenInClassPostgres{db: db}
}