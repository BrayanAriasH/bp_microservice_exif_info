package util

import (
	"bytes"
	"image/jpeg"
	"log"

	"github.com/disintegration/imaging"
)

func CompressImage(originalImage []byte, width int, heigth int) (compressedImage []byte, err error) {
	reader := bytes.NewReader(originalImage)
	img, err := jpeg.Decode(reader)
	if err != nil {
		log.Printf("Error in CompressImage: %s", err)
		return nil, err
	}
	compressedImg := imaging.Resize(img, width, heigth, imaging.Lanczos)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, compressedImg, nil)
	if err != nil {
		log.Printf("Error in CompressImage: %s", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

func CreateCompressedImage(originalImage []byte) (compressedImage []byte, err error) {
	return CompressImage(originalImage, 800, 0)
}
