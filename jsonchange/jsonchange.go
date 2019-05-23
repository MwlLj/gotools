package jsonchange

import (
	"encoding/json"
)

func JsonToByByte(input []byte, inObj interface{}, outObj interface{}, callback func(in interface{}, out interface{})) ([]byte, error) {
	err := json.Unmarshal(input, inObj)
	if err != nil {
		return nil, err
	}
	callback(inObj, outObj)
	return json.Marshal(outObj)
}
