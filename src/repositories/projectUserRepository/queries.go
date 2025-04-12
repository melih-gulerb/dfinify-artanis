package projectUserRepository

import (
	"divine-shield/src/models"
	"fmt"
)

var RegisterProjectUserQuery = func(user models.ProjectUser) string {
	return fmt.Sprintf("INSERT INTO dbo.ProjectUsers (Id, RoleId, ProjectId, UserId) VALUES ('%s', '%s', '%s', '%s')", user.Id, user.RoleId, user.ProjectId, user.UserId)
}
