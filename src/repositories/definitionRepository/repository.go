package definitionRepository

import (
	"artanis/src/logging"
	"artanis/src/models/entities"
	servicemodal "artanis/src/models/services"
	"database/sql"
)

type DefinitionRepository struct {
	DB *sql.DB
}

func NewDefinitionRepository(db *sql.DB) *DefinitionRepository {
	return &DefinitionRepository{DB: db}
}

func (repo *DefinitionRepository) RegisterDefinition(definition entities.Definition) error {
	_, err := repo.DB.Exec(RegisterDefinitionQuery(),
		sql.Named("Id", definition.Id),
		sql.Named("Name", definition.Name),
		sql.Named("Value", definition.Value),
		sql.Named("CollectionId", definition.CollectionId))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *DefinitionRepository) PaginateDefinitions(collectionId string, limit, offset int) ([]entities.Definition, error) {
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

	var definitions []entities.Definition
	for rows.Next() {
		var definition entities.Definition
		err := rows.Scan(&definition.Id, &definition.Name, &definition.Value, &definition.CreatedAt)
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

func (repo *DefinitionRepository) UpdateDefinitionName(id string, name string) error {
	_, err := repo.DB.Exec(UpdateDefinitionNameQuery(),
		sql.Named("Id", id),
		sql.Named("Name", name))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *DefinitionRepository) UpdateDefinitionValue(id string, value string) error {
	_, err := repo.DB.Exec(UpdateDefinitionValueQuery(),
		sql.Named("Id", id),
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

func (repo *DefinitionRepository) GetDefinition(id string) *entities.Definition {
	var definition entities.Definition

	err := repo.DB.QueryRow(GetDefinitionByIdQuery(), sql.Named("Id", id)).Scan(&definition.Id, &definition.Value,
		&definition.Name, &definition.CollectionId, &definition.CreatedAt)

	if err != nil {
		return nil
	}

	return &definition
}

func (repo *DefinitionRepository) GetDefinitionDetail(definitionId string) *servicemodal.DefinitionDetail {
	var definitionDetail servicemodal.DefinitionDetail
	err := repo.DB.QueryRow(GetDefinitionDetail(), sql.Named("Id", definitionId)).Scan(&definitionDetail.CollectionName,
		&definitionDetail.DefinitionName, &definitionDetail.ProjectId, &definitionDetail.ProjectName, &definitionDetail.OldValue)

	if err != nil {
		return nil
	}

	return &definitionDetail
}
