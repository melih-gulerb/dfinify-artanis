package definitionChangeRepository

import (
	"artanis/src/models/requests"
	"database/sql"
	"github.com/google/uuid"
)

type DefinitionChangeRepository struct {
	DB *sql.DB
}

func NewDefinitionChangeRepository(db *sql.DB) *DefinitionChangeRepository {
	return &DefinitionChangeRepository{DB: db}
}

func (repo *DefinitionChangeRepository) RegisterDefinitionChange(change requests.RegisterDefinitionChange) error {
	_, err := repo.DB.Exec(RegisterDefinitionChange(), sql.Named("Id", uuid.New().String()), sql.Named("UserId", change.UserId),
		sql.Named("DefinitionId", change.DefinitionId), sql.Named("OldValue", change.OldValue), sql.Named("NewValue", change.NewValue))
	return err
}

func (repo *DefinitionChangeRepository) UpdateDefinitionChangeState(change requests.UpdateDefinitionChange) error {
	_, err := repo.DB.Exec(UpdateDefinitionChangeState(), sql.Named("DefinitionId", change.DefinitionId), sql.Named("State", change.State))
	return err
}

func (repo *DefinitionChangeRepository) GetDefinitionChange(definitionId string) (string, error) {
	var value string
	err := repo.DB.QueryRow(GetDefinitionChangeState(), sql.Named("DefinitionId", definitionId)).Scan(&value)
	if err != nil {
		return "", err
	}

	return value, err
}
