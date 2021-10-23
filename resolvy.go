package resolvy

import (
	"encoding/json"
)

// TagName is the default tag name used by resolvy.
const TagName = "resolvy"

// MarshalJSON is a utility for producing JSON
// from the marshalled representation of v.
func MarshalJSON(v interface{}, config MarshalConfig) ([]byte, error) {
	res, err := Marshal(v, config)
	if err != nil {
		return nil, err
	}
	return json.Marshal(res)
}
