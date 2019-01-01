package file

import (
	"bufio"
	"os"
)

func ReadBinary(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != err {
		return nil, err
	}
	defer file.Close()
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	var size int64 = stats.Size()
	bytes := make([]byte, size)
	buf := bufio.NewReader(file)
	_, err = buf.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
