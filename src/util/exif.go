package util

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

var exposureModeMap = map[int]string{
	0: "auto_exposure",
	1: "manual",
	2: "auto_bracket",
}

var orientation = map[int]string{
	0: "horizontal",
	1: "vertical",
}

func GetExifStringDataByTag(exif *exif.Exif, fieldName exif.FieldName) (value string, err error) {
	valueRaw, err := exif.Get(fieldName)
	if err != nil {
		return "", err
	}
	value = valueRaw.String()
	value = strings.ReplaceAll(value, "\"", "")
	return value, nil
}

func GetExifIntDataByTag(exif *exif.Exif, fieldName exif.FieldName) (value int, err error) {
	valueRaw, err := exif.Get(fieldName)
	if err != nil {
		return 0, err
	}
	result, err := valueRaw.Int(0)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetExifUIntDataByTag(exif *exif.Exif, fieldName exif.FieldName) (value uint, err error) {
	valueRaw, err := exif.Get(fieldName)
	if err != nil {
		return 0, err
	}
	result, err := valueRaw.Int(0)
	if err != nil {
		return 0, err
	}
	return uint(result), nil
}

func GetExifFloatDataByTag(exif *exif.Exif, fieldName exif.FieldName) (value float64, err error) {
	valueRaw, err := exif.Get(fieldName)
	if err != nil {
		return 0, err
	}
	num, den, err := valueRaw.Rat2(0)
	if err != nil {
		return 0, err
	}
	return float64(num) / float64(den), nil
}

func GetExifFromBytes(data []byte) (x *exif.Exif, err error) {
	x, err = exif.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return x, err
}

func GetStringExposureMode(x *exif.Exif) (value string, err error) {
	i, err := GetExifIntDataByTag(x, exif.ExposureMode)
	if err != nil {
		return "", err
	}
	return exposureModeMap[i], nil
}

func GetStringOrientation(x *exif.Exif) (value string, err error) {
	resolutionX, err := GetExifIntDataByTag(x, exif.PixelXDimension)
	if err != nil {
		return "", err
	}
	resolutionY, err := GetExifIntDataByTag(x, exif.PixelYDimension)
	if err != nil {
		return "", err
	}
	i := 0
	fmt.Println("resolutionX", resolutionX, "resolutionY", resolutionY)
	if resolutionX >= resolutionY {
		i = 0
	} else {
		i = 1
	}
	return orientation[i], nil
}
