package projectUserRepository

import (
	"artanis/src/models/entities"
	"artanis/src/models/enums"
	"database/sql"
)

type ProjectUserRepository struct {
	DB *sql.DB
}

func NewProjectUserRepository(db *sql.DB) *ProjectUserRepository {
	return &ProjectUserRepository{DB: db}
}

func (repo *ProjectUserRepository) RegisterProjectUser(user entities.ProjectUser) error {
	_, err := repo.DB.Exec(RegisterProjectUserQuery())
	return err
}

func (repo *ProjectUserRepository) GetProjectUser(userId, projectId string) *enums.ProjectRole {
	var role enums.ProjectRole
	err := repo.DB.QueryRow(GetProjectUserQuery(), sql.Named("UserId", userId),
		sql.Named("ProjectId", projectId)).Scan(&role)

	if err != nil {
		return nil
	}

	return &role
}
