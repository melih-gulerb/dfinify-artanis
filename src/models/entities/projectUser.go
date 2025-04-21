package entities

import "time"

type ProjectUser struct {
	Id             string
	ProjectId      string
	UserId         string
	RoleId         string
	SlackChannelId string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
