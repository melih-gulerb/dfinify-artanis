package requests

import "artanis/src/models/enums"

type RegisterDefinitionChange struct {
	UserId       string
	DefinitionId string
	OldValue     string
	NewValue     string
}

type UpdateDefinitionChange struct {
	DefinitionId string                      `json:"definitionId"`
	State        enums.DefinitionChangeState `json:"state"`
}
