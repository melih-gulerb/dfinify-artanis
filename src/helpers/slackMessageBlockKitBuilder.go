package helpers

import (
	helpermodal "artanis/src/models/helpers"
	"fmt"
	"github.com/slack-go/slack"
)

func CreateDefinitionChangeRequestSlackBlocks(request helpermodal.CreateDefinitionChangeRequestSlackModel) []slack.Block {
	section1 := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: "*:memo: Definition change request*",
		},
		nil,
		nil,
	)

	section2 := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: fmt.Sprintf("Requester - *%s*", request.UserName),
		},
		nil,
		nil,
	)

	section3 := slack.NewSectionBlock(
		nil,
		[]*slack.TextBlockObject{
			{
				Type: slack.MarkdownType,
				Text: fmt.Sprintf("*Project*\n%s", request.ProjectName),
			},
			{
				Type: slack.MarkdownType,
				Text: fmt.Sprintf("*Collection*\n%s", request.CollectionName),
			},
			{
				Type: slack.MarkdownType,
				Text: fmt.Sprintf("*Definition*\n%s", request.DefinitionName),
			},
		},
		nil,
	)

	section4 := slack.NewSectionBlock(
		nil,
		[]*slack.TextBlockObject{
			{
				Type: slack.MarkdownType,
				Text: fmt.Sprintf("*Old Value*\n%s \n *New Value*\n%s", request.OldValue, request.NewValue),
			},
		},
		nil,
	)

	divider := slack.NewDividerBlock()

	section5 := slack.NewActionBlock(
		"approve_deny_action",
		slack.NewButtonBlockElement(
			"approve_button",
			fmt.Sprintf("approve:%s", request.DefinitionId),
			&slack.TextBlockObject{
				Type:  slack.PlainTextType,
				Text:  "Approve",
				Emoji: true,
			},
		).WithStyle(slack.StylePrimary),
		slack.NewButtonBlockElement(
			"deny_button",
			fmt.Sprintf("deny:%s", request.DefinitionId),
			&slack.TextBlockObject{
				Type:  slack.PlainTextType,
				Text:  "Deny",
				Emoji: true,
			},
		).WithStyle(slack.StyleDanger),
	)

	blocks := []slack.Block{section1, section2, section3, section4, divider, section5}
	return blocks
}
