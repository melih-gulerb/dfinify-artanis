package projectRepository

var RegisterProjectQuery = func() string {
	return `
INSERT INTO dbo.Projects (Id, Name, Description) 
VALUES (@Id, @Name, @Description)
`
}

var PaginateProjectsQuery = func() string {
	return `
SELECT Id, Name, Description 
FROM dbo.Projects 
WHERE OrganizationId = @OrganizationId 
  AND DeletedAt IS NULL
GROUP BY OrganizationId, Id, Name, Description 
ORDER BY CreatedAt DESC 
OFFSET @Offset ROWS 
FETCH NEXT @Limit ROWS ONLY
`
}

var UpdateProjectQuery = func() string {
	return `
UPDATE dbo.Projects 
SET Name = @Name, Description = @Description
WHERE Id = @Id
`
}

var DeleteProjectQuery = func() string {
	return `
UPDATE dbo.Projects 
SET DeletedAt = GETDATE() 
WHERE Id = @Id
`
}

var GetDashboardCountsQuery = func(organizationId string) string {
	return `
SELECT
    p.OrganizationId,
    COUNT(DISTINCT p.Id) AS ProjectCount,
    COUNT(c.Id) AS CollectionCount,
    COUNT(d.Id) AS DefinitionCount
FROM
    Projects p
LEFT JOIN
    Collections c ON p.Id = c.ProjectId
LEFT JOIN
    Definitions d ON c.Id = d.CollectionId
WHERE
    p.OrganizationId = '` + organizationId + `'
GROUP BY
    p.OrganizationId;`
}
