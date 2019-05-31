package ospath

import (
	"os"
)

func Exists(path *string) bool {
	_, err := os.Stat(*path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDirsIfNotExists(path *string) error {
	if Exists(path) == false {
		err := os.MkdirAll(*path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
