package jsonchange

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func JsonToByByte(input []byte, inObj interface{}, outObj interface{}, callback func(in interface{}, out interface{})) ([]byte, error) {
	err := json.Unmarshal(input, inObj)
	if err != nil {
		return nil, err
	}
	callback(inObj, outObj)
	return json.Marshal(outObj)
}

func JsonToByIoReader(reader io.Reader, inObj interface{}, outObj interface{}, callback func(in interface{}, out interface{})) ([]byte, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return JsonToByByte(b, inObj, outObj, callback)
}
