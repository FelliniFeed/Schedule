package models

type Timepair struct {
	Id        string `json:"Id" db:"Id"`
	StartPair string `json:"StartPair" db:"StartPair"`
	EndPair   string `json:"EndPair" db:"EndPair"`
}