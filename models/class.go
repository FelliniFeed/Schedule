package models

import "github.com/google/uuid"

type Class struct {
	Id uuid.UUID `json:"Id" db:"Id"`
	Name string `json:"Name" db:"Name"`
}