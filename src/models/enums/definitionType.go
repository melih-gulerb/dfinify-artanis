package enums

type DefinitionType int

const (
	String DefinitionType = iota
	Int
	Float
	Boolean
	Date
)

func (r DefinitionType) String() string {
	switch r {
	case String:
		return "String"
	case Int:
		return "Int"
	case Float:
		return "Float"
	case Boolean:
		return "Boolean"
	case Date:
		return "Date"
	default:
		return "Unknown"
	}
}
