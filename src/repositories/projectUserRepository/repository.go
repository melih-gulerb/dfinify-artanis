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

func (repo *ProjectUserRepository) GetProjectAdminsForSlackUser(projectId string) []string {
	var slackChannelIds []string

	rows, err := repo.DB.Query(GetProjectAdminsForSlackUserQuery(), sql.Named("ProjectId", projectId))
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var slackChannelId string
		if err := rows.Scan(&slackChannelId); err != nil {
			continue
		}
		slackChannelIds = append(slackChannelIds, slackChannelId)
	}

	if err = rows.Err(); err != nil {
		return nil
	}

	return slackChannelIds
}
