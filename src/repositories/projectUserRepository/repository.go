package projectUserRepository

import (
	"artanis/src/models/entities"
	"artanis/src/models/enums"
	"artanis/src/models/responses"
	"database/sql"
)

type ProjectUserRepository struct {
	DB *sql.DB
}

func NewProjectUserRepository(db *sql.DB) *ProjectUserRepository {
	return &ProjectUserRepository{DB: db}
}

func (repo *ProjectUserRepository) RegisterProjectUser(user entities.ProjectUser) error {
	_, err := repo.DB.Exec(RegisterProjectUserQuery(),
		sql.Named("UserId", user.UserId),
		sql.Named("ProjectId", user.ProjectId),
		sql.Named("Role", user.Role))

	return err
}

func (repo *ProjectUserRepository) UpdateProjectUserRole(projectUserId string, role enums.ProjectRole) error {
	_, err := repo.DB.Exec(UpdateProjectUserRoleQuery(),
		sql.Named("Id", projectUserId),
		sql.Named("Role", role))

	return err
}

func (repo *ProjectUserRepository) DeleteProjectUser(projectUserId string) error {
	_, err := repo.DB.Exec(DeleteProjectUserRoleQuery(),
		sql.Named("Id", projectUserId))

	return err
}

func (repo *ProjectUserRepository) Paginate(projectId string, limit, offset int) []responses.ProjectUserPaginationResponse {
	var projectUsers []responses.ProjectUserPaginationResponse

	rows, err := repo.DB.Query(PaginateProjectUsersQuery(), sql.Named("ProjectId", projectId),
		sql.Named("Limit", limit), sql.Named("Offset", offset))
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var projectUser responses.ProjectUserPaginationResponse
		if err := rows.Scan(&projectUser.Id, &projectUser.Role, &projectUser.UserId, &projectUser.CreatedAt); err != nil {
			continue
		}
		projectUsers = append(projectUsers, projectUser)
	}

	if err = rows.Err(); err != nil {
		return nil
	}

	return projectUsers
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
