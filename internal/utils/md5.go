package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func CalcMd5(f *multipart.FileHeader) (string, error) {
	r, err := f.Open()
	defer r.Close()
	if err != nil {
		return "", err
	}
	h := md5.New()
	_, err = io.Copy(h, r)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
