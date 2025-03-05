package requests

type RegisterProject struct {
	Name        string
	Description string
}

type UpdateProject struct {
	Id          string
	Name        string
	Description string
}
