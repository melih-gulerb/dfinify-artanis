package queries

import (
	"artanis/src/models"
	"fmt"
	"strings"
)

var RegisterProjectQuery = func(project models.Project) string {
	return fmt.Sprintf("INSERT INTO dbo.Projects (Id, Name, Description) VALUES ('%s', '%s', '%s')", project.Id, project.Name, project.Description)
}

var PaginateProjectsQuery = func(organizationId string, limit, offset int) string {
	return fmt.Sprintf("SELECT Id, Name, Description FROM dbo.Projects WHERE OrganizationId = '%s' AND DeletedAt IS NULL "+
		"GROUP BY OrganizationId, Id, Name, Description ORDER BY CreatedAt DESC OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", organizationId, offset, limit)
}

var UpdateProjectQuery = func(id string, name string, description string) string {
	query := "UPDATE dbo.Projects SET "
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

var DeleteProjectQuery = func(id string) string {
	return fmt.Sprintf("UPDATE dbo.Projects SET DeletedAt = GETDATE() WHERE Id = '%s'", id)
}
