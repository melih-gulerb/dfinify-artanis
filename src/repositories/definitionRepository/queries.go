package definitionRepository

var RegisterDefinitionQuery = func() string {
	return `
INSERT INTO dbo.Definitions (Id, Name, Value, Description) 
VALUES (@Id, @Name, @Value, @Description)
`
}

var PaginateDefinitionsQuery = func() string {
	return `
SELECT Id, Name, Description 
FROM dbo.Definitions 
WHERE CollectionId = @CollectionId 
  AND DeletedAt IS NULL
GROUP BY CollectionId, Id, Name, Description 
ORDER BY CreatedAt DESC 
OFFSET @Offset ROWS 
FETCH NEXT @Limit ROWS ONLY
`
}

var UpdateDefinitionQuery = func() string {
	return `
UPDATE dbo.Definitions 
SET Name = @Name, Value = @Value
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
