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
