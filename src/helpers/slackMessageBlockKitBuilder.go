package helpers

import (
	helpermodal "artanis/src/models/helpers"
	"fmt"
	"github.com/slack-go/slack"
	"strings"
)

func truncateText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	return text[:maxLength] + "...\n(content truncated due to length)"
}

func createFormattedDiffSummary(oldValue, newValue string) string {
	oldLines := len(strings.Split(oldValue, "\n"))
	newLines := len(strings.Split(newValue, "\n"))

	lineDiff := newLines - oldLines

	var diffSummary string
	if lineDiff > 0 {
		diffSummary = fmt.Sprintf("%d lines added", lineDiff)
	} else if lineDiff < 0 {
		diffSummary = fmt.Sprintf("%d lines removed", -lineDiff)
	} else {
		diffSummary = "Same number of lines, content modified"
	}

	return diffSummary
}

func CreateDefinitionChangeRequestSlackBlocks(request helpermodal.CreateDefinitionChangeRequestSlackModel) []slack.Block {
	header := slack.NewHeaderBlock(
		&slack.TextBlockObject{
			Type:  slack.PlainTextType,
			Text:  "Definition Change Request",
			Emoji: true,
		},
	)

	context := slack.NewContextBlock(
		"requester_context",
		slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf(":bust_in_silhouette: *Requested by:* %s", request.UserMail), false, false),
	)

	infoSection := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: ":information_source: *Change Details*",
		},
		[]*slack.TextBlockObject{
			{
				Type: slack.MarkdownType,
				Text: fmt.Sprintf("*Project:*\n%s", request.ProjectName),
			},
			{
				Type: slack.MarkdownType,
				Text: fmt.Sprintf("*Collection:*\n%s", request.CollectionName),
			},
			{
				Type: slack.MarkdownType,
				Text: fmt.Sprintf("*Definition:*\n%s", request.DefinitionName),
			},
		},
		nil,
	)

	divider1 := slack.NewDividerBlock()

	diffSummary := createFormattedDiffSummary(request.OldValue, request.NewValue)

	changesSummary := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: fmt.Sprintf(":arrows_counterclockwise: *Changes Summary:* %s", diffSummary),
		},
		nil,
		nil,
	)

	const previewLength = 600
	oldValuePreview := truncateText(request.OldValue, previewLength)
	newValuePreview := truncateText(request.NewValue, previewLength)

	previewSection := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: "*Preview:*",
		},
		nil,
		nil,
	)

	oldValueSection := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: "*Old Value Preview:*\n```\n" + oldValuePreview + "\n```",
		},
		nil,
		nil,
	)

	newValueSection := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: "*New Value Preview:*\n```\n" + newValuePreview + "\n```",
		},
		nil,
		nil,
	)

	viewFullChangesButton := slack.NewAccessory(
		slack.NewButtonBlockElement(
			"view_full_diff",
			fmt.Sprintf("view_diff:%s", request.DefinitionId),
			&slack.TextBlockObject{
				Type:  slack.PlainTextType,
				Text:  "View Full Diff",
				Emoji: true,
			},
		),
	)

	viewFullChangesSection := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: ":mag: *Need more details?*",
		},
		nil,
		viewFullChangesButton,
	)

	divider2 := slack.NewDividerBlock()

	actionPrompt := slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: ":point_right: *Please review this change request:*",
		},
		nil,
		nil,
	)

	actionButtons := slack.NewActionBlock(
		"approve_deny_action",
		slack.NewButtonBlockElement(
			"approve_button",
			fmt.Sprintf("approve:%s", request.DefinitionId),
			&slack.TextBlockObject{
				Type:  slack.PlainTextType,
				Text:  "Approve Change",
				Emoji: true,
			},
		).WithStyle(slack.StylePrimary),
		slack.NewButtonBlockElement(
			"deny_button",
			fmt.Sprintf("deny:%s", request.DefinitionId),
			&slack.TextBlockObject{
				Type:  slack.PlainTextType,
				Text:  "Deny Change",
				Emoji: true,
			},
		).WithStyle(slack.StyleDanger),
	)

	blocks := []slack.Block{
		header,
		context,
		divider1,
		infoSection,
		changesSummary,
		previewSection,
		oldValueSection,
		newValueSection,
		viewFullChangesSection,
		divider2,
		actionPrompt,
		actionButtons,
	}

	return blocks
}
