package resolvy

import (
	"encoding/json"
)

// TagName is the default tag name used by resolvy.
const TagName = "resolvy"

func MarshalJSON(v interface{}, config MarshalConfig) ([]byte, error) {
	res, err := Marshal(v, config)
	if err != nil {
		return nil, err
	}
	return json.Marshal(res)
}
