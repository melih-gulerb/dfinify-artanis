package projectRepository

import (
	"artanis/src/logging"
	"artanis/src/models"
	"artanis/src/models/responses"
	"database/sql"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (repo *ProjectRepository) RegisterProject(project models.Project) error {
	_, err := repo.DB.Exec(RegisterProjectQuery(),
		sql.Named("Id", project.Id),
		sql.Named("Name", project.Name),
		sql.Named("Description", project.Description))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *ProjectRepository) PaginateProjects(organizationId string, limit, offset int) ([]models.Project, error) {
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

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.Id, &project.Name, &project.Description)
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
