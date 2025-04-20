package clients

import (
	"context"
	"fmt"
	"github.com/slack-go/slack"
)

type Slack struct {
	client *slack.Client
}

func NewSlackClient(slackToken string) *Slack {
	return &Slack{
		client: slack.New(slackToken),
	}
}

type SlackMessage struct {
	Blocks []slack.Block `json:"blocks"`
}

func (s *Slack) SendBlockKitMessage(channelID string, blocks []slack.Block) error {
	_, _, err := s.client.PostMessageContext(
		context.Background(),
		channelID,
		slack.MsgOptionBlocks(blocks...),
	)
	if err != nil {
		return fmt.Errorf("failed to send BlockKit message: %w", err)
	}
	return nil
}

func (s *Slack) UpdateBlockKitMessage(channelID, timestamp string, blocks []slack.Block) error {
	_, _, _, err := s.client.UpdateMessage(
		channelID,
		timestamp,
		slack.MsgOptionBlocks(blocks...),
	)
	if err != nil {
		return fmt.Errorf("failed to update BlockKit message: %w", err)
	}
	return nil
}
