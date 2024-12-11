package models

import "github.com/google/uuid"

type Subject struct {
	Id uuid.UUID `json:"Id" db:"Id"`
	Name string `json:"Name" db:"Name"`
}