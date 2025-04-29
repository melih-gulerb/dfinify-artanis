package helpers

import (
	"artanis/src/models/enums"
	"errors"
	"strconv"
	"time"
)

func GetTypedValue(definitionType enums.DefinitionType, value string) (interface{}, error) {
	switch definitionType {
	case enums.String:
		return value, nil
	case enums.Int:
		return strconv.ParseInt(value, 10, 64)
	case enums.Float:
		return strconv.ParseFloat(value, 64)
	case enums.Boolean:
		return strconv.ParseBool(value)
	case enums.Date:
		return time.Parse(time.RFC3339, value)
	default:
		return nil, errors.New("unsupported data type")
	}
}
