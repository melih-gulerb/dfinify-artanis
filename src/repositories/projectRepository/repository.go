package projectRepository

import (
	"artanis/src/logging"
	"artanis/src/models"
	"database/sql"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (repo *ProjectRepository) RegisterProject(project models.Project) error {
	_, err := repo.DB.Exec(RegisterProjectQuery(project))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *ProjectRepository) PaginateProjects(organizationId string, limit, offset int) ([]models.Project, error) {
	query := PaginateProjectsQuery(organizationId, limit, offset)
	rows, err := repo.DB.Query(query)
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
	_, err := repo.DB.Exec(UpdateProjectQuery(id, name, description))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *ProjectRepository) DeleteProject(id string) error {
	_, err := repo.DB.Exec(DeleteProjectQuery(id))
	if err != nil {
	}

	return err
}
