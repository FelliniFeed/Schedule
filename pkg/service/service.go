package service

import (
	"schedule/models"
	"schedule/pkg/repository"

	"github.com/google/uuid"
)

type Class interface {
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

type Service struct {
	Class
	StudentInClass
	Student
	Schedule
	Timepair
	Teacher
	Subject
}

func NewService(repos *repository.Reposiroty) *Service {
	return &Service{
		Class: NewClassService(repos.Class),
		StudentInClass: NewStudenInClassService(repos.StudentInClass),
		Student: NewStudentService(repos.Student),
		Schedule: NewScheduleService(repos.Schedule),
		Timepair: NewTimepairService(repos.Timepair),
		Teacher: NewTeacherService(repos.Teacher),
		Subject: NewSubjectService(repos.Subject),
	}
}
