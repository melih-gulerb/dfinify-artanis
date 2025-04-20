package services

import (
	"artanis/src/clients"
	"artanis/src/helpers"
	helpers2 "artanis/src/models/helpers"
	"artanis/src/models/requests"
	models "artanis/src/models/services"
	"artanis/src/repositories/definitionChangeRepository"
)

type DefinitionChangeService struct {
	dcb   definitionChangeRepository.DefinitionChangeRepository
	slack *clients.Slack
}

func NewDefinitionChangeService(dcb *definitionChangeRepository.DefinitionChangeRepository, slack *clients.Slack) *DefinitionChangeService {
	return &DefinitionChangeService{dcb: dcb, slack: slack}
}

func (s *DefinitionChangeService) Register(request models.RegisterDefinitionChange) {
	change := requests.RegisterDefinitionChange{
		DefinitionId: request.DefinitionId,
		UserId:       request.UserId,
		OldValue:     request.OldValue,
		NewValue:     request.NewValue,
	}

	err := s.dcb.RegisterDefinitionChange(change)
	if err != nil {
		return
	}

	s.SendToSlack("", request)
}

func (s *DefinitionChangeService) SendToSlack(slackChannelId string, model helpers2.CreateDefinitionChangeRequestSlackModel) error {
	s.slack.SendBlockKitMessage(slackChannelId, helpers.CreateDefinitionChangeRequestSlackBlocks(model))
	return nil
}

func SendToMail() error {
	return nil
}
