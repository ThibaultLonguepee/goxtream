package dtos

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Duration int

func (ps *Duration) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch x := v.(type) {
	case nil:
		*ps = 0
	case int:
		*ps = Duration(x)
	case string:
		value, err := strconv.Atoi(x)
		if err != nil {
			return fmt.Errorf("duration contains a non-decimal string")
		}
		*ps = Duration(value)
	case float64:
		*ps = Duration(int(x))
	default:
		return fmt.Errorf("paths unsupported type: %T", x)
	}

	return nil
}
