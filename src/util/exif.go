package util

import (
	"bytes"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

var exposureModeMap = map[int]string{
	0: "auto_exposure",
	1: "manual",
	2: "auto_bracket",
}

var orientation = map[int]string{
	1: "horizontal",
	2: "horizontal",
	3: "horizontal",
	4: "horizontal",
	5: "vertical",
	6: "vertical",
	7: "vertical",
	8: "vertical",
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
	i, err := GetExifIntDataByTag(x, exif.Orientation)
	if err != nil {
		return "", err
	}
	return orientation[i], nil
}
