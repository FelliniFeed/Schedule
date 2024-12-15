package repository

import (
	"schedule/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Class interface{
	GetClasses() ([]models.Class, error)
	CreateClass(class models.Class) (uuid.UUID, error)
}
type StudentInClass interface{}
type Student interface{
	GetStudents() ([]models.Student, error)
	CreateStudent(student models.Student) (uuid.UUID, error)
}
type Schedule interface{}
type Timepair interface{}
type Teacher interface{}
type Subject interface{}	

type Reposiroty struct {
	Class
	StudentInClass
	Student
	Schedule
	Timepair
	Teacher
	Subject
}

func NewRepository(db *sqlx.DB) *Reposiroty {
	return &Reposiroty{
		Class: NewClassPostgres(db),
		StudentInClass: NewStudenInClassPostgres(db),
		Student: NewStudenPostgres(db),
		Schedule: NewSchedulePostgres(db),
		Timepair: NewTimepairPostgres(db),
		Teacher: NewTeacherPostgres(db),
		Subject: NewSubjectPostgres(db),
	}
}