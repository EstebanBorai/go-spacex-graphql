package util

import (
	"encoding/json"
	"io"
)

// JSONStruct fills the entity interface with the JSON content
// if the fields match
func JSONStruct(body io.ReadCloser, entity interface{}) error {
	return json.NewDecoder(body).Decode(entity)
}
