package projectUserRepository

var RegisterProjectUserQuery = func() string {
	return `
INSERT INTO dbo.ProjectUsers (Id, RoleId, ProjectId, UserId) 
VALUES (@Id, @RoleId, @ProjectId, @UserId)
`
}

var GetProjectUserQuery = func() string {
	return `
	SELECT Role FROM dbo.ProjectUsers
	WHERE ProjectId = @ProjectId AND UserId = @UserId AND DeletedAt IS NULL
`
}

var GetProjectAdminsForSlackUserQuery = func() string {
	return `
	SELECT SlackChannelId FROM dbo.ProjectUsers
	WHERE ProjectId = @ProjectId AND Role = 1 AND DeletedAt IS NULL
`
}
