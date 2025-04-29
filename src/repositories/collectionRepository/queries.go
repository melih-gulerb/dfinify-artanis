package collectionRepository

var RegisterCollectionQuery = func() string {
	return `
INSERT INTO dbo.Collections (Id, Name, Description, ProjectId) 
VALUES (@Id, @Name, @Description, @ProjectId)
`
}

var PaginateCollectionsQuery = func() string {
	return `
SELECT Id, Name, Description, CreatedAt
FROM dbo.Collections 
WHERE ProjectId = @ProjectId 
  AND DeletedAt IS NULL
ORDER BY CreatedAt DESC 
OFFSET @Offset ROWS 
FETCH NEXT @Limit ROWS ONLY
`
}

var UpdateCollectionQuery = func() string {
	return `
UPDATE dbo.Collections 
SET Name = @Name, Description = @Description
WHERE Id = @Id
`
}

var DeleteCollectionQuery = func() string {
	return `
UPDATE dbo.Collections 
SET DeletedAt = GETDATE() 
WHERE Id = @Id
`
}
