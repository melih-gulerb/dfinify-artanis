package responses

import (
	"artanis/src/models/enums"
	"time"
)

type ProjectUserPaginationResponse struct {
	Id        string            `json:"id"`
	UserId    string            `json:"userId"`
	UserMail  string            `json:"userMail"`
	Username  string            `json:"username"`
	Role      enums.ProjectRole `json:"role"`
	CreatedAt time.Time         `json:"createdAt"`
}
