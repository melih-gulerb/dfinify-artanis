package repositories

import (
	"artanis/src/logging"
	"artanis/src/models"
	"artanis/src/repositories/queries"
	"database/sql"
)

type DefinitionRepository struct {
	DB *sql.DB
}

func NewDefinitionRepository(db *sql.DB) *DefinitionRepository {
	return &DefinitionRepository{DB: db}
}

func (repo *DefinitionRepository) RegisterDefinition(Definition models.Definition) error {
	_, err := repo.DB.Exec(queries.RegisterDefinitionQuery(Definition))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *DefinitionRepository) PaginateDefinitions(collectionId string, limit, offset int) ([]models.Definition, error) {
	query := queries.PaginateDefinitionsQuery(collectionId, limit, offset)
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
	_, err := repo.DB.Exec(queries.UpdateDefinitionQuery(id, name, value))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *DefinitionRepository) DeleteDefinition(id string) error {
	_, err := repo.DB.Exec(queries.DeleteDefinitionQuery(id))
	if err != nil {
	}

	return err
}
