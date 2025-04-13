package projectUserRepository

var RegisterProjectUserQuery = func() string {
	return `
INSERT INTO dbo.ProjectUsers (Id, RoleId, ProjectId, UserId) 
VALUES (@Id, @RoleId, @ProjectId, @UserId)
`
}
