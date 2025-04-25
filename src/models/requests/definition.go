package requests

type RegisterDefinition struct {
	Name         string `json:"name"`
	Value        string `json:"value"`
	CollectionId string `json:"collectionId"`
	ProjectId    string `json:"projectId"`
}

type UpdateDefinition struct {
	Id             string
	Name           string `json:"name"`
	Value          string `json:"value"`
	ProjectId      string `json:"projectId"`
	ProjectName    string `json:"projectName"`
	CollectionName string `json:"collectionName"`
}
type UpdateDefinitionName struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

type UpdateDefinitionValue struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}
