package enums

type DefinitionChangeState int

const (
	Submitted DefinitionChangeState = iota
	ChangeApproved
	Rejected
)

func (r DefinitionChangeState) String() string {
	switch r {
	case Submitted:
		return "Submitted"
	case ChangeApproved:
		return "Approved"
	case Rejected:
		return "Rejected"
	default:
		return "Unknown"
	}
}
