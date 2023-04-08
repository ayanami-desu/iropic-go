package optimize

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ayanami-desu/iropic-go/internal/bootstrap"
	"github.com/ayanami-desu/iropic-go/internal/utils"
	"github.com/chai2010/webp"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"strings"
)

func Conv2Webp(f *multipart.FileHeader, timePath string, fileName string, fileType string) error {
	var err error
	var buf bytes.Buffer
	var img image.Image
	data, err := f.Open()
	if err != nil {
		return err
	}

	// Decode pic
	if strings.Contains(fileType, "jpeg") || strings.Contains(fileType, "jpg") {
		img, err = jpeg.Decode(data)
	} else if strings.Contains(fileType, "png") {
		img, err = png.Decode(data)
	} else {
		return errors.New("不支持的文件类型")
	}
	if err != nil {
		return err
	}

	// Encode webp
	err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 90})
	if err != nil {
		return err
	}

	fullPath := fmt.Sprintf("%s%s%s", bootstrap.Conf.PathPrefix, bootstrap.WebpImgPath, timePath)
	err = utils.CheckPath(fullPath)
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath + fileName + ".webp")
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
