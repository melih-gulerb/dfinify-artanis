package projectRepository

import (
	"artanis/src/logging"
	"artanis/src/models/entities"
	"artanis/src/models/responses"
	"database/sql"
	"errors"
	"time"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (repo *ProjectRepository) RegisterProject(project entities.Project) error {
	_, err := repo.DB.Exec(RegisterProjectQuery(),
		sql.Named("Id", project.Id),
		sql.Named("Name", project.Name),
		sql.Named("Description", project.Description),
		sql.Named("OrganizationId", project.OrganizationId))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *ProjectRepository) PaginateProjects(organizationId string, limit, offset int) ([]entities.Project, error) {
	rows, err := repo.DB.Query(PaginateProjectsQuery(),
		sql.Named("OrganizationId", organizationId),
		sql.Named("Limit", limit),
		sql.Named("Offset", offset))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logging.Log(logging.ERROR, err.Error())
		}
	}(rows)

	var projects []entities.Project
	for rows.Next() {
		var project entities.Project
		var createdAt string
		err := rows.Scan(&project.Id, &project.Name, &project.Description, &createdAt)
		project.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			logging.Log(logging.ERROR, err.Error())
			return nil, err
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		logging.Log(logging.ERROR, err.Error())
		return nil, err
	}

	return projects, nil
}

func (repo *ProjectRepository) UpdateProject(id string, name string, description string) error {
	_, err := repo.DB.Exec(UpdateProjectQuery(),
		sql.Named("Id", id),
		sql.Named("Name", name),
		sql.Named("Description", description))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *ProjectRepository) DeleteProject(id string) error {
	_, err := repo.DB.Exec(DeleteProjectQuery(), sql.Named("Id", id))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *ProjectRepository) GetDashboardCounts(organizationId string) (responses.DashboardResponse, error) {
	var dashboard responses.DashboardResponse

	rows, err := repo.DB.Query(GetDashboardCountsQuery(organizationId))
	if err != nil {
		return dashboard, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&dashboard.ProjectCount, &dashboard.CollectionCount, &dashboard.DefinitionCount)
		if err != nil {
			return dashboard, err
		}
	}

	return dashboard, nil
}

func (repo *ProjectRepository) GetProjectFeed(projectId string) ([]responses.ProjectFeed, error) {
	var projectFeeds []responses.ProjectFeed

	rows, err := repo.DB.Query(GetProjectFeed(), sql.Named("ProjectId", projectId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var projectFeed responses.ProjectFeed
		err := rows.Scan(&projectFeed.CollectionName, &projectFeed.DefinitionId, &projectFeed.DefinitionValue)
		if err != nil {
			logging.Log(logging.ERROR, err.Error())
			return nil, err
		}
		projectFeeds = append(projectFeeds, projectFeed)
	}

	return projectFeeds, nil
}

func (repo *ProjectRepository) ValidateSecret(projectId, secret string) error {
	var foundProjectId string
	err := repo.DB.QueryRow(ValidateSecret(), sql.Named("Id", projectId), sql.Named("Secret", secret)).Scan(&foundProjectId)
	if err != nil {
		return err
	}
	if foundProjectId != projectId {
		return errors.New("invalid secret")
	}

	return nil
}

func (repo *ProjectRepository) UpdateProjectSecret(id, secret string) error {
	_, err := repo.DB.Exec(UpdateSecret(),
		sql.Named("Id", id),
		sql.Named("Secret", secret))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}
