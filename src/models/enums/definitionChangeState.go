package enums

type DefinitionChangeState int

const (
	Submitted DefinitionChangeState = iota
	ChangeApproved
	ChangeRejected
)

func (r DefinitionChangeState) String() string {
	switch r {
	case Submitted:
		return "Submitted"
	case ChangeApproved:
		return "Approved"
	case ChangeRejected:
		return "ChangeRejected"
	default:
		return "Unknown"
	}
}
