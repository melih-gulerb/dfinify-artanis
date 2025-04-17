package requests

type RegisterCollection struct {
	Name        string
	Description string
	ProjectId   string
}

type UpdateCollection struct {
	Id          string
	ProjectId   string
	Name        string
	Description string
}
