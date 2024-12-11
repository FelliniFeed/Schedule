package repository

import (
	"schedule/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClassPostgres struct {
	db *sqlx.DB
}

func NewClassPostgres(db *sqlx.DB) *ClassPostgres {
	return &ClassPostgres{db:db}
}

func (r *ClassPostgres) GetClasses() ([]models.Class, error) {
	var class []models.Class
	query := `SELECT * FROM class`
	err := r.db.Select(&class, query)

	return class, err
}

func (r *ClassPostgres) CreateClass(class models.Class) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO class (name) VALUES ($1) RETURNING id;`
	row := r.db.QueryRow(query, class.Name)
	err := row.Scan(&id)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}