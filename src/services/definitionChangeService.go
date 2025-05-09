package services

import (
	"artanis/src/clients"
	"artanis/src/helpers"
	"artanis/src/models/enums"
	helperModel "artanis/src/models/helpers"
	"artanis/src/models/requests"
	models "artanis/src/models/services"
	"artanis/src/repositories/definitionChangeRepository"
	"fmt"
)

type DefinitionChangeService struct {
	dcb   *definitionChangeRepository.DefinitionChangeRepository
	slack *clients.Slack
}

func NewDefinitionChangeService(dcb *definitionChangeRepository.DefinitionChangeRepository, slack *clients.Slack) *DefinitionChangeService {
	return &DefinitionChangeService{dcb: dcb, slack: slack}
}

func (s *DefinitionChangeService) Register(request models.RegisterDefinitionChange) error {
	change := requests.RegisterDefinitionChange{
		DefinitionId: request.DefinitionId,
		UserId:       request.UserId,
		OldValue:     request.OldValue,
		NewValue:     request.NewValue,
	}

	err := s.dcb.RegisterDefinitionChange(change)
	if err != nil {
		return err
	}

	blockKit := helperModel.CreateDefinitionChangeRequestSlackModel{
		ProjectName:    request.ProjectName,
		CollectionName: request.CollectionName,
		DefinitionId:   request.DefinitionId,
		DefinitionName: request.DefinitionName,
		OldValue:       request.OldValue,
		NewValue:       request.NewValue,
		UserMail:       request.UserMail,
	}

	for _, id := range request.SlackChannelIds {
		err := s.sendToSlack(id, blockKit)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func (s *DefinitionChangeService) UpdateState(definitionId string, state enums.DefinitionChangeState) error {
	err := s.dcb.UpdateDefinitionChangeState(requests.UpdateDefinitionChange{
		DefinitionId: definitionId,
		State:        state,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *DefinitionChangeService) GetDefinitionChange(definitionId string) (string, error) {
	value, err := s.dcb.GetDefinitionChange(definitionId)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (s *DefinitionChangeService) sendToSlack(slackChannelId string, model helperModel.CreateDefinitionChangeRequestSlackModel) error {
	err := s.slack.SendBlockKitMessage(slackChannelId, helpers.CreateDefinitionChangeRequestSlackBlocks(model))
	if err != nil {
		return err
	}

	return nil
}
