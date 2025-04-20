package collectionRepository

import (
	"artanis/src/logging"
	"artanis/src/models/entities"
	"database/sql"
)

type CollectionRepository struct {
	DB *sql.DB
}

func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{DB: db}
}

func (repo *CollectionRepository) RegisterCollection(collection entities.Collection) error {
	_, err := repo.DB.Exec(RegisterCollectionQuery(),
		sql.Named("Id", collection.Id),
		sql.Named("Name", collection.Name),
		sql.Named("Description", collection.Description),
		sql.Named("ProjectId", collection.ProjectId))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}
	return err
}

func (repo *CollectionRepository) PaginateCollections(projectId string, limit, offset int) ([]entities.Collection, error) {
	rows, err := repo.DB.Query(PaginateCollectionsQuery(),
		sql.Named("ProjectId", projectId),
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

	var collections []entities.Collection
	for rows.Next() {
		var collection entities.Collection
		err := rows.Scan(&collection.Id, &collection.Name, &collection.Description, &collection.CreatedAt)
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
	_, err := repo.DB.Exec(UpdateCollectionQuery(),
		sql.Named("Id", id),
		sql.Named("Name", name),
		sql.Named("Description", description))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}
	return err
}

func (repo *CollectionRepository) DeleteCollection(id string) error {
	_, err := repo.DB.Exec(DeleteCollectionQuery(), sql.Named("Id", id))
	if err != nil {
		logging.Log(logging.ERROR, err.Error())
	}
	return err
}
