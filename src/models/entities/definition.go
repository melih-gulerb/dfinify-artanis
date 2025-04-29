package entities

import (
	"artanis/src/models/enums"
	"time"
)

type Definition struct {
	Id           string
	Name         string
	Value        string
	Type         enums.DefinitionType
	CollectionId string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
