package models

type CreateDefinitionChangeRequestSlackModel struct {
	ProjectName, CollectionName, DefinitionId, DefinitionName, OldValue, NewValue, UserName, UserMail string
}
