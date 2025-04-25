package service_modal

type RegisterDefinitionChange struct {
	DefinitionId    string
	OldValue        string
	NewValue        string
	ProjectName     string
	CollectionName  string
	DefinitionName  string
	UserMail        string
	UserId          string
	SlackChannelIds []string
}
