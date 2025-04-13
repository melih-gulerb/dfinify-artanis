package clients

import "artanis/src/models/enums"

type User struct {
	Id               string
	Name             string
	Email            string
	Password         string
	State            enums.UserState
	OrganizationId   string
	OrganizationRole enums.OrganizationRole
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}

type AuthResponse struct {
	User User `json:"user"`
}

type AuthRequest struct {
	Token string `json:"token"`
}
