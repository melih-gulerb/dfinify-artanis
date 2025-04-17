package services

import (
	"artanis/src/models/requests"
	"artanis/src/repositories/definitionChangeRepository"
	"artanis/src/repositories/definitionRepository"
)

type DefinitionChangeService struct {
	dcb definitionChangeRepository.DefinitionChangeRepository
	db  definitionRepository.DefinitionRepository
}

func NewDefinitionChangeService(dcb definitionChangeRepository.DefinitionChangeRepository, db definitionRepository.DefinitionRepository) *DefinitionChangeService {
	return &DefinitionChangeService{dcb: dcb, db: db}
}

func (s *DefinitionChangeService) Register(definitionId, userId, newValue string) {
	definition := s.db.GetDefinition(definitionId)

	if definition == nil {
		return
	}

	change := requests.RegisterDefinitionChange{
		DefinitionId: definitionId,
		UserId:       userId,
		OldValue:     definition.Value,
		NewValue:     newValue,
	}

	err := s.dcb.RegisterDefinitionChange(change)
	if err != nil {
		return
	}
}

func SendToSlack() error {
	return nil
}

func SendToMail() error {
	return nil
}
