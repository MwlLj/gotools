package datachange

import (
	"fmt"
	"testing"
)

func TestVector2MapStringVector(t *testing.T) {
	type Test struct {
		Uuid string
		Name string
	}
	data := []Test{}
	data = append(data, Test{"1", "a"})
	data = append(data, Test{"1", "b"})
	data = append(data, Test{"1", "c"})
	data = append(data, Test{"2", "a"})
	data = append(data, Test{"2", "b"})
	result := Vector2MapStringVector(&data, func(d interface{}) *string {
		uuid := d.(Test).Uuid
		return &uuid
	})
	if result != nil {
		for key, value := range *result {
			vec := []Test{}
			for _, item := range value {
				vec = append(vec, item.(Test))
			}
			fmt.Println(key, vec)
		}
	}
}
