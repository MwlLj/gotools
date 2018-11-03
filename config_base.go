package configs

import (
	"io/ioutil"
	"os"
)

type CConfigBase struct {
}

func (this *CConfigBase) Load(path string, defaultString string) (string, error) {
	if this.eixsts(path) {
		text, err := this.readAll(path)
		if err != nil {
			return "", err
		} else {
			return text, nil
		}
	} else {
		err := this.writeString(path, defaultString)
		if err != nil {
			return "", err
		}
		return defaultString, nil
	}
}

func (*CConfigBase) eixsts(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (*CConfigBase) readAll(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (*CConfigBase) writeString(path string, text string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, e := f.WriteString(text)
	if e != nil {
		return e
	}
	return nil
}
