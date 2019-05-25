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

func JsonToByByteWithRet(input []byte, inObj interface{}, outObj interface{}, callback func(in interface{}, out interface{}) error) ([]byte, error) {
	err := json.Unmarshal(input, inObj)
	if err != nil {
		return nil, err
	}
	err = callback(inObj, outObj)
	if err != nil {
		return nil, err
	}
	return json.Marshal(outObj)
}

func JsonToByIoReader(reader io.Reader, inObj interface{}, outObj interface{}, callback func(in interface{}, out interface{})) ([]byte, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return JsonToByByte(b, inObj, outObj, callback)
}

func JsonToString(inObj interface{}, out *string) error {
	b, err := json.Marshal(&inObj)
	if err != nil {
		return err
	}
	*out = string(b)
	return nil
}

func JsonFromByte(b []byte, outObj interface{}) error {
	return json.Unmarshal(b, &outObj)
}

func JsonFromString(s *string, outObj interface{}) error {
	return JsonFromByte([]byte(*s), outObj)
}
