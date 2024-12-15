package repository

import (
	"schedule/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StudenPostgres struct {
	db *sqlx.DB
}

func NewStudenPostgres(db *sqlx.DB) *StudenPostgres{
	return &StudenPostgres{db:db}
}

func (r *StudenPostgres) GetStudents() ([]models.Student, error) {
	var students []models.Student
	query := `SELECT * FROM student`
	err := r.db.Select(&students, query)
	return students, err
}

func (r *StudenPostgres) CreateStudent(student models.Student) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO student (first_name, last_name, patronymic, birth_day, address) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	row := r.db.QueryRow(query, student.FirstName, student.LastName, student.Patronymic, student.BirthDay, student.Address)
	err := row.Scan(&id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}