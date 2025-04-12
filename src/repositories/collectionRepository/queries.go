package collectionRepository

import (
	"artanis/src/models"
	"fmt"
	"strings"
)

var RegisterCollectionQuery = func(collection models.Collection) string {
	return fmt.Sprintf("INSERT INTO dbo.Collections (Id, Name, Description) VALUES ('%s', '%s', '%s')", collection.Id, collection.Name, collection.Description)
}

var PaginateCollectionsQuery = func(projectId string, limit, offset int) string {
	return fmt.Sprintf("SELECT Id, Name, Description FROM dbo.Collections WHERE ProjectId = '%s' AND DeletedAt IS NULL "+
		"GROUP BY ProjectId, Id, Name, Description ORDER BY CreatedAt DESC OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", projectId, offset, limit)
}

var UpdateCollectionQuery = func(id string, name string, description string) string {
	query := "UPDATE dbo.Collections SET "
	var updates []string

	if name != "" {
		updates = append(updates, fmt.Sprintf("Name = '%s'", name))
	}
	if description != "" {
		updates = append(updates, fmt.Sprintf("Description = '%s'", description))
	}

	query += fmt.Sprintf("%s WHERE Id = '%s'", strings.Join(updates, ", "), id)
	return query
}

var DeleteCollectionQuery = func(id string) string {
	return fmt.Sprintf("UPDATE dbo.Collections SET DeletedAt = GETDATE() WHERE Id = '%s'", id)
}
