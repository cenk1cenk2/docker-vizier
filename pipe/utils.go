package pipe

import (
	"encoding/json"
	"fmt"
	"time"
)

type StringDuration struct {
	time.Duration
}

func (duration *StringDuration) UnmarshalJSON(b []byte) error {
	var unmarshalledJson interface{}

	err := json.Unmarshal(b, &unmarshalledJson)
	if err != nil {
		return err
	}

	switch value := unmarshalledJson.(type) {
	case float64:
		duration.Duration = time.Duration(value)
	case string:
		duration.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid duration: %#v", unmarshalledJson)
	}

	return nil
}
