package enums

type ProjectRole int

const (
	ProjectOwner ProjectRole = iota
	ProjectUser
)

func (r ProjectRole) String() string {
	switch r {
	case ProjectOwner:
		return "ProjectOwner"
	case ProjectUser:
		return "ProjectUser"
	default:
		return "Unknown"
	}
}
