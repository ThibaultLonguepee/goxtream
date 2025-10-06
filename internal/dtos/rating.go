package dtos

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Rating float64

func (f *Rating) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch x := v.(type) {
	case nil:
		*f = 0.0
	case string:
		if x == "" {
			*f = 0.0
		} else {
			val, err := strconv.ParseFloat(x, 64)
			if err != nil {
				return err
			}
			*f = Rating(val)
		}
	case float64:
		*f = Rating(x)
	case int:
		*f = Rating(float64(x))
	default:
		return fmt.Errorf("unsupported type: %T", x)
	}

	return nil
}
