package models

import "artanis/src/models/enums"

type DefinitionChange struct {
	Id           string
	DefinitionId string
	UserId       string
	OldValue     string
	NewValue     string
	CreatedAt    string
	UpdatedAt    string
	State        enums.DefinitionChangeState
}
