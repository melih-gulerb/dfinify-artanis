package responses

import "time"

type PaginateCollectionResponse struct {
	CollectionResponse []CollectionResponse `json:"items"`
	TotalCount         int                  `json:"totalCount"`
}

type CollectionResponse struct {
	Id          string
	Name        string
	Description string
	CreatedAt   time.Time
}
