package queries

import (
	"artanis/src/models"
	"fmt"
	"strings"
)

var RegisterDefinitionQuery = func(definition models.Definition) string {
	return fmt.Sprintf("INSERT INTO dbo.Definitions (Id, Name, Description) VALUES ('%s', '%s', '%s')", definition.Id, definition.Name, definition.Value)
}

var PaginateDefinitionsQuery = func(projectId string, limit, offset int) string {
	return fmt.Sprintf("SELECT Id, Name, Description FROM dbo.Definitions WHERE ProjectId = '%s' AND DeletedAt IS NULL "+
		"GROUP BY ProjectId, Id, Name, Description ORDER BY CreatedAt DESC OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", projectId, offset, limit)
}

var UpdateDefinitionQuery = func(id string, name string, value string) string {
	query := "UPDATE dbo.Definitions SET "
	var updates []string

	if name != "" {
		updates = append(updates, fmt.Sprintf("Name = '%s'", name))
	}
	if value != "" {
		updates = append(updates, fmt.Sprintf("Value = '%s'", value))
	}

	query += fmt.Sprintf("%s WHERE Id = '%s'", strings.Join(updates, ", "), id)
	return query
}

var DeleteDefinitionQuery = func(id string) string {
	return fmt.Sprintf("UPDATE dbo.Definitions SET DeletedAt = GETDATE() WHERE Id = '%s'", id)
}
