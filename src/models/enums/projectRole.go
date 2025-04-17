package enums

type ProjectRole int

const (
	ProjectUser ProjectRole = iota
	ProjectOwner
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
