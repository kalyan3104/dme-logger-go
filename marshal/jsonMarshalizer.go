package marshal

import (
	"encoding/json"
)

// JSONMarshalizer -
type JSONMarshalizer struct {
}

// Marshal does the actual serialization of an object
func (marshalizer *JSONMarshalizer) Marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// Unmarshal does the actual deserialization of an object
func (marshalizer *JSONMarshalizer) Unmarshal(obj interface{}, buff []byte) error {
	return json.Unmarshal(buff, obj)
}

// IsInterfaceNil returns true if there is no value under the interface
func (marshalizer *JSONMarshalizer) IsInterfaceNil() bool {
	return marshalizer == nil
}
