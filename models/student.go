package models

import "github.com/google/uuid"

type Student struct {
	Id         uuid.UUID `json:"Id" db:"Id"`
	FirstName  string    `json:"FirstName" db:"FirstName"`
	LastName   string    `json:"LastName" db:"LastName"`
	Patronymic string    `json:"Patronymic" db:"Patronymic"`
	BirthDay   string    `json:"BirthDay" db:"BirthDay"`
	Address    string    `json:"Address" db:"Address"`
}