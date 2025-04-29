package client_modal

import "artanis/src/models/enums"

type User struct {
	Id               string                 `json:"id"`
	Name             string                 `json:"name"`
	Email            string                 `json:"email"`
	Password         string                 `json:"password"`
	State            enums.UserState        `json:"state"`
	OrganizationId   string                 `json:"organizationId"`
	OrganizationRole enums.OrganizationRole `json:"organizationRole"`
	SlackChannelId   string                 `json:"slackChannelId"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	DeletedAt        string                 `json:"deletedAt"`
}

type UserInformation struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserInformationResponse struct {
	Success         bool              `json:"success"`
	Message         string            `json:"message"`
	UserInformation []UserInformation `json:"data"`
}

type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    User   `json:"data"`
}

type AuthRequest struct {
	Token string `json:"token"`
}
type UserInformationRequest struct {
	UserIds []string `json:"userIds"`
}
