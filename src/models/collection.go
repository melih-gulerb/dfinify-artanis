package models

import "time"

type Collection struct {
	Id          string
	Name        string
	Description string
	ProjectId   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
