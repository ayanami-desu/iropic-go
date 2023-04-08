package utils

import (
	"errors"
	"os"
)

func CheckFile(path string) (*os.File, error) {
	var f *os.File
	fileInfo, err := os.Stat(path)
	if err != nil {
		f, err = os.Create(path)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	if fileInfo.IsDir() {
		return nil, errors.New("此路径是目录")
	}
	f, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func CheckPath(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	}
	if !fileInfo.IsDir() {
		return errors.New("此路径不是目录")
	}
	return nil
}
