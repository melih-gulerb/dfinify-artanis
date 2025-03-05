package responses

import "time"

type PaginateDefinitionResponse struct {
	DefinitionResponse []DefinitionResponse `json:"items"`
	TotalCount         int                  `json:"totalCount"`
}

type DefinitionResponse struct {
	Id        string
	Name      string
	Value     string
	CreatedAt time.Time
}
