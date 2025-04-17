package definitionRepository

import (
	"artanis/src/logging"
	"artanis/src/models"
	"database/sql"
)

type DefinitionRepository struct {
	DB *sql.DB
}

func NewDefinitionRepository(db *sql.DB) *DefinitionRepository {
	return &DefinitionRepository{DB: db}
}

func (repo *DefinitionRepository) RegisterDefinition(definition models.Definition) error {
	_, err := repo.DB.Exec(RegisterDefinitionQuery(),
		sql.Named("Id", definition.Id),
		sql.Named("Name", definition.Name),
		sql.Named("Value", definition.Value))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *DefinitionRepository) PaginateDefinitions(collectionId string, limit, offset int) ([]models.Definition, error) {
	rows, err := repo.DB.Query(PaginateDefinitionsQuery(),
		sql.Named("CollectionId", collectionId),
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

	var definitions []models.Definition
	for rows.Next() {
		var definition models.Definition
		err := rows.Scan(&definition.Id, &definition.Name, &definition.Value)
		if err != nil {
			logging.Log(logging.ERROR, err.Error())
			return nil, err
		}
		definitions = append(definitions, definition)
	}

	if err = rows.Err(); err != nil {
		logging.Log(logging.ERROR, err.Error())
		return nil, err
	}

	return definitions, nil
}

func (repo *DefinitionRepository) UpdateDefinition(id string, name string, value string) error {
	_, err := repo.DB.Exec(UpdateDefinitionQuery(),
		sql.Named("Id", id),
		sql.Named("Name", name),
		sql.Named("Value", value))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *DefinitionRepository) DeleteDefinition(id string) error {
	_, err := repo.DB.Exec(DeleteDefinitionQuery(), sql.Named("Id", id))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *DefinitionRepository) GetDefinition(id string) *models.Definition {
	var definition models.Definition

	err := repo.DB.QueryRow(GetDefinitionByIdQuery(), sql.Named("Id", id)).Scan(&definition.Id, &definition.Value,
		&definition.Name, &definition.CollectionId)

	if err != nil {
		return nil
	}

	return &definition
}
