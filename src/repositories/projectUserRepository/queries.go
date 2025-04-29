package projectUserRepository

var RegisterProjectUserQuery = func() string {
	return `
INSERT INTO dbo.ProjectUsers (Id, Role, ProjectId, UserId) 
VALUES (@Id, @Role, @ProjectId, @UserId)
`
}

var UpdateProjectUserRoleQuery = func() string {
	return `
UPDATE dbo.ProjectUsers 
SET Role = @Role
WHERE Id = @Id
`
}

var DeleteProjectUserRoleQuery = func() string {
	return `
UPDATE dbo.ProjectUsers 
SET DeletedAt = GETDATE()
WHERE Id = @Id
`
}

var PaginateProjectUsersQuery = func() string {
	return `
SELECT Id, Role, UserId, CreatedAt
FROM dbo.ProjectUsers 
WHERE ProjectId = @ProjectId 
  AND DeletedAt IS NULL
ORDER BY CreatedAt DESC 
OFFSET @Offset ROWS 
FETCH NEXT @Limit ROWS ONLY
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
