package requests

type RegisterProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateProject struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
