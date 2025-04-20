package models

type RegisterDefinitionChange struct {
	DefinitionId   string
	ProjectId      string
	OldValue       string
	NewValue       string
	ProjectName    string
	CollectionName string
	DefinitionName string
	UserName       string
	UserMail       string
	UserId         string
}
