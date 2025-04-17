package projectUserRepository

import (
	"artanis/src/models"
	"artanis/src/models/enums"
	"database/sql"
)

type ProjectUserRepository struct {
	DB *sql.DB
}

func NewProjectUserRepository(db *sql.DB) *ProjectUserRepository {
	return &ProjectUserRepository{DB: db}
}

func (repo *ProjectUserRepository) RegisterProjectUser(user models.ProjectUser) error {
	_, err := repo.DB.Exec(RegisterProjectUserQuery())
	return err
}

func (repo *ProjectUserRepository) GetProjectUser(userId, projectId string) *enums.ProjectRole {
	var role *enums.ProjectRole
	_ = repo.DB.QueryRow(RegisterProjectUserQuery(), sql.Named("UserId", userId),
		sql.Named("ProjectId", projectId)).Scan(&role)

	return role
}
