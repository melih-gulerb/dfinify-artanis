package responses

import (
	"artanis/src/models/enums"
	"time"
)

type ProjectUserPaginationResponse struct {
	Id        string            `json:"id"`
	UserId    string            `json:"userId"`
	Role      enums.ProjectRole `json:"role"`
	CreatedAt time.Time         `json:"createdAt"`
}
