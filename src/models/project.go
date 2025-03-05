package models

import "time"

type Project struct {
	Id             string
	Name           string
	Description    string
	OrganizationId string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
