package resolvy

import (
	"encoding/json"
)

const TagName = "resolvy"

func MarshalJSON(v interface{}, config MarshalConfig) ([]byte, error) {
	res, err := Marshal(v, config)
	if err != nil {
		return nil, err
	}
	return json.Marshal(res)
}
