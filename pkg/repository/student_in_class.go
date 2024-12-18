package repository

import (
	"errors"
	"schedule/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StudenInClassPostgres struct {
	db *sqlx.DB
}

func NewStudenInClassPostgres(db *sqlx.DB) *StudenInClassPostgres {
	return &StudenInClassPostgres{db: db}
}

func (r *StudenInClassPostgres) GetStudentInClass(studentId, classId string) ([]models.StudentInClass, error) {

	var studentInClass []models.StudentInClass
	var args []interface{}
	
	query := `SELECT * FROM student_class WHERE`

	if studentId != "" {
		query += ` studentId = $1`
		args = append(args, studentId)
	}

	if classId != "" {
		if len(args) > 0 {
			query += ` AND`
		}
		query += ` classId = $2`
		args = append(args, classId)
	}

	if len(args) == 0 {
		return nil, errors.New("no parameters provided")
	}

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()	
	
	for rows.Next() {
		var student models.StudentInClass
		if err := rows.Scan(&studentInClass); err != nil {
			return nil, err
		}
		studentInClass = append(studentInClass, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return studentInClass, nil
}

func (r *StudenInClassPostgres) CreateStudentInClass(studentInClass models.StudentInClass) (uuid.UUID, error) {
	//TODO: нарушает ограничение внешнего ключа \"student_class_student_id_fkey я даун помогите
	var id uuid.UUID
	query := `INSERT INTO student_class (student_id, class_id) VALUES ($1, $2) RETURNING id;`
	row := r.db.QueryRow(query, studentInClass.Student, studentInClass.Class)
	err := row.Scan(&id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}