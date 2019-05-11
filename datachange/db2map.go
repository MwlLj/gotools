package datachange

import (
	"log"
	"reflect"
)

func Vector2MapStringVector(data interface{}, keyFn func(data interface{}) *string) *map[string][]interface{} {
	if pointer := reflect.ValueOf(data); pointer.Kind() == reflect.Ptr {
		if slice := reflect.ValueOf(pointer.Elem().Interface()); slice.Kind() == reflect.Slice {
			retMap := map[string][]interface{}{}
			for i := 0; i < slice.Len(); i++ {
				item := slice.Index(i)
				inter := item.Interface()
				key := keyFn(inter)
				if key == nil {
					continue
				}
				if _, ok := retMap[*key]; ok {
					// exist
					retMap[*key] = append(retMap[*key], inter)
				} else {
					vec := []interface{}{}
					vec = append(vec, inter)
					retMap[*key] = vec
				}
			}
			return &retMap
		}
	} else {
		log.Fatalln("[datachange.Vector2MapStringVector error] please input *[]interface{} type")
	}
	return nil
}
