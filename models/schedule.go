package models

import "github.com/google/uuid"

type Shedule struct {
	Id    uuid.UUID `json:"Id" db:"Id"`
	Date  string    `json:"Date" db:"Date"`
	Class uuid.UUID `json:"Class" db:"Class"`
	NumberPair int    `json:"NumberPair" db:"NumberPair"`
	Teacher uuid.UUID `json:"Teacher" db:"Teacher"`
	Subject uuid.UUID `json:"Subject" db:"Subject"`
	ClassRoom string    `json:"ClassRoom" db:"ClassRoom"`
}