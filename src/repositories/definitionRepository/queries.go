package definitionRepository

var RegisterDefinitionQuery = func() string {
	return `
INSERT INTO dbo.Definitions (Id, Name, Value, CollectionId) 
VALUES (@Id, @Name, @Value, @CollectionId)
`
}

var PaginateDefinitionsQuery = func() string {
	return `
SELECT Id, Name, Value, CreatedAt
FROM dbo.Definitions 
WHERE CollectionId = @CollectionId 
  AND DeletedAt IS NULL
ORDER BY CreatedAt DESC 
OFFSET @Offset ROWS 
FETCH NEXT @Limit ROWS ONLY
`
}

var UpdateDefinitionNameQuery = func() string {
	return `
UPDATE dbo.Definitions 
SET Name = @Name
WHERE Id = @Id
`
}
var UpdateDefinitionValueQuery = func() string {
	return `
UPDATE dbo.Definitions 
SET Value = @Value, UpdatedAt = GETDATE()
WHERE Id = @Id
`
}

var DeleteDefinitionQuery = func() string {
	return `
UPDATE dbo.Definitions 
SET DeletedAt = GETDATE() 
WHERE Id = @Id
`
}

var GetDefinitionByIdQuery = func() string {
	return `
SELECT Id, Name, Value, CollectionId, CreatedAt
FROM dbo.Definitions 
WHERE Id = @Id 
  AND DeletedAt IS NULL
`
}

var GetDefinitionDetail = func() string {
	return `
SELECT c.Name AS CollectionName,
       d.Name As DefinitionName,
       p.Id AS ProjectId, p.Name AS ProjectName,
       d.Value As OldValue
FROM dbo.Definitions d
JOIN dbo.Collections c ON d.CollectionId = c.Id AND c.DeletedAt IS NULL
JOIN dbo.Projects p ON c.ProjectId = p.Id AND p.DeletedAt IS NULL
WHERE d.Id = @Id 
  AND d.DeletedAt IS NULL
`
}
