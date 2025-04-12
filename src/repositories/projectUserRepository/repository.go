package projectUserRepository

import (
	"database/sql"
	"divine-shield/src/models"
)

type ProjectUserRepository struct {
	DB *sql.DB
}

func NewProjectUserRepository(db *sql.DB) *ProjectUserRepository {
	return &ProjectUserRepository{DB: db}
}

func (repo *ProjectUserRepository) RegisterProjectUser(user models.ProjectUser) error {
	_, err := repo.DB.Exec(RegisterProjectUserQuery(user))
	return err
}
