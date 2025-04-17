package requests

type RegisterDefinition struct {
	Name         string
	Value        string
	CollectionId string
	ProjectId    string
}

type UpdateDefinition struct {
	Id        string
	Name      string
	Value     string
	ProjectId string
}
