package models

import "github.com/google/uuid"

type StudentInClass struct {
	Id    uuid.UUID `json:"Id" db:"Id"`
	Class uuid.UUID `json:"Class" db:"Class"`
	Student uuid.UUID `json:"Student" db:"Student"`

}