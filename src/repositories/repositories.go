package repositories

import (
	"artanis/src/repositories/collectionRepository"
	"artanis/src/repositories/definitionChangeRepository"
	"artanis/src/repositories/definitionRepository"
	"artanis/src/repositories/projectRepository"
	"artanis/src/repositories/projectUserRepository"
)

type Repositories struct {
	ProjectRepository          *projectRepository.ProjectRepository
	ProjectUserRepository      *projectUserRepository.ProjectUserRepository
	CollectionRepository       *collectionRepository.CollectionRepository
	DefinitionRepository       *definitionRepository.DefinitionRepository
	DefinitionChangeRepository *definitionChangeRepository.DefinitionChangeRepository
}
