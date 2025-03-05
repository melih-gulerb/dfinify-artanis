package repositories

import (
	"artanis/src/logging"
	"artanis/src/models"
	"artanis/src/repositories/queries"
	"database/sql"
)

type CollectionRepository struct {
	DB *sql.DB
}

func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{DB: db}
}

func (repo *CollectionRepository) RegisterCollection(collection models.Collection) error {
	_, err := repo.DB.Exec(queries.RegisterCollectionQuery(collection))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *CollectionRepository) PaginateCollections(projectId string, limit, offset int) ([]models.Collection, error) {
	query := queries.PaginateCollectionsQuery(projectId, limit, offset)
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

	var collections []models.Collection
	for rows.Next() {
		var collection models.Collection
		err := rows.Scan(&collection.Id, &collection.Name, &collection.Description)
		if err != nil {
			logging.Log(logging.ERROR, err.Error())
			return nil, err
		}
		collections = append(collections, collection)
	}

	if err = rows.Err(); err != nil {
		logging.Log(logging.ERROR, err.Error())
		return nil, err
	}

	return collections, nil
}

func (repo *CollectionRepository) UpdateCollection(id string, name string, description string) error {
	_, err := repo.DB.Exec(queries.UpdateCollectionQuery(id, name, description))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}

	return err
}

func (repo *CollectionRepository) DeleteCollection(id string) error {
	_, err := repo.DB.Exec(queries.DeleteCollectionQuery(id))
	if err != nil {
	}

	return err
}
