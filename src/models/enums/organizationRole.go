package enums

type OrganizationRole int

const (
	OrganizationOwner OrganizationRole = iota
	OrganizationUser
	PersonalUser
)

func (r OrganizationRole) String() string {
	switch r {
	case OrganizationOwner:
		return "OrganizationOwner"
	case OrganizationUser:
		return "OrganizationUser"
	case PersonalUser:
		return "PersonalUser"
	default:
		return "Unknown"
	}
}

func (r OrganizationRole) IsOrganizationUser() bool {
	return r == OrganizationUser
}

func (r OrganizationRole) IsValidOrganizationRole(role OrganizationRole) bool {
	return role == OrganizationOwner || role == OrganizationUser || role == PersonalUser
}
