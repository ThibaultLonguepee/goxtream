package dtos

import (
	"encoding/json"
	"fmt"
)

type Paths []string

func (ps *Paths) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch x := v.(type) {
	case nil:
		*ps = nil
	case string:
		*ps = make(Paths, 1)
		(*ps)[0] = string(data)
	case []interface{}:
		*ps = make(Paths, 0)
		for _, el := range x {
			if s, ok := el.(string); ok {
				*ps = append(*ps, s)
			} else {
				return fmt.Errorf("paths array contains non-string element: %T", el)
			}
		}
	default:
		return fmt.Errorf("paths unsupported type: %T", x)
	}

	return nil
}
