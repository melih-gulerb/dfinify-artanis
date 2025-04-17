package clients

import "artanis/src/models/enums"

type User struct {
	Id               string                 `json:"id"`
	Name             string                 `json:"name"`
	Email            string                 `json:"email"`
	Password         string                 `json:"password"`
	State            enums.UserState        `json:"state"`
	OrganizationId   string                 `json:"organizationId"`
	OrganizationRole enums.OrganizationRole `json:"organizationRole"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	DeletedAt        string                 `json:"deletedAt"`
}

type AuthResponse struct {
	User User `json:"user"`
}

type AuthRequest struct {
	Token string `json:"token"`
}
