package entities

import (
	"artanis/src/models/enums"
	"time"
)

type ProjectUser struct {
	Id             string
	ProjectId      string
	UserId         string
	Role           enums.ProjectRole
	SlackChannelId string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
