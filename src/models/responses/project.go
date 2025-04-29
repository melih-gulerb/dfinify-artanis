package responses

import "time"

type PaginateProjectResponse struct {
	ProjectResponse []ProjectResponse `json:"items"`
	TotalCount      int               `json:"totalCount"`
}

type ProjectResponse struct {
	Id          string
	Name        string
	Description string
	CreatedAt   time.Time
}

type DashboardResponse struct {
	ProjectCount    int
	CollectionCount int
	DefinitionCount int
}

type ProjectFeed struct {
	CollectionName  string
	DefinitionId    string
	DefinitionValue string
}

type DefinitionFeed struct {
	DefinitionId    string `json:"definitionId"`
	DefinitionValue string `json:"definitionValue"`
}
