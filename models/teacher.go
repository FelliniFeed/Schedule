package models

import "github.com/google/uuid"

type Teacher struct {
	Id         uuid.UUID `json:"Id" db:"Id"`
	FirstName  string    `json:"FirstName" db:"FirstName"`
	LastName   string    `json:"LastName" db:"LastName"`
	Patronymic string    `json:"Patronymic" db:"Patronymic"`
}