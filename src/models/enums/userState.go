package enums

type UserState int

const (
	Approved UserState = iota
	WaitingOrganizationOwnerApprove
	Deleted
)

func (r UserState) String() string {
	switch r {
	case Approved:
		return "Approved"
	case WaitingOrganizationOwnerApprove:
		return "WaitingOrganizationOwnerApprove"
	case Deleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}
