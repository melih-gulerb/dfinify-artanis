package entities

import "time"

type Definition struct {
	Id           string
	Name         string
	Value        string
	CollectionId string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
