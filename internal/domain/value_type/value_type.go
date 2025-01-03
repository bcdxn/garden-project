package value_type

import (
	"encoding/json"
	"time"
)

type NullableTime struct {
	Time   time.Time
	IsNull bool
}

func (t *NullableTime) MarshalJSON() ([]byte, error) {
	if t.IsNull {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time)
}
