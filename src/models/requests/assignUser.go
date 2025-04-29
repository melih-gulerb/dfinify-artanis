package requests

import "artanis/src/models/enums"

type Register struct {
	UserId    string
	ProjectId string
	Role      enums.ProjectRole
}

type UpdateRole struct {
	Role enums.ProjectRole
}
