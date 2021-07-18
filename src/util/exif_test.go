package util

import (
	"io/ioutil"
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
	"github.com/rwcarlsen/goexif/exif"
)

func TestGetStringOrientationVertical(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/images/DSC03783.JPG")
	x, err := GetExifFromBytes(file)
	if err != nil {
		t.Errorf("Error on TestGetStringOrientation %s", err.Error())
	}
	orientation, err := GetStringOrientation(x)
	if err != nil {
		t.Errorf("Error on TestGetStringOrientation %s", err.Error())
	}
	assert.Equals("vertical", orientation, t)
}

func TestGetStringOrientationHorizontal(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/images/DSC04249.JPG")
	x, err := GetExifFromBytes(file)
	if err != nil {
		t.Errorf("Error on TestGetStringOrientation %s", err.Error())
	}
	orientation, err := GetStringOrientation(x)
	if err != nil {
		t.Errorf("Error on TestGetStringOrientation %s", err.Error())
	}
	assert.Equals("horizontal", orientation, t)
}

func TestGetStringOrientation(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/images/DSC04249.JPG")
	x, err := GetExifFromBytes(file)
	if err != nil {
		t.Errorf("Error on TestGetStringOrientation %s", err.Error())
	}
	expected := "DSC-HX400V"
	result, err := GetExifStringDataByTag(x, exif.Model)
	if err != nil {
		t.Errorf("Error on TestGetStringOrientation %s", err.Error())
	}
	assert.NotNil(result, t)
	assert.Equals(expected, result, t)
}

func TestGetExifUIntDataByTag(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/images/DSC04249.JPG")
	x, err := GetExifFromBytes(file)
	if err != nil {
		t.Errorf("Error on TestGetExifUIntDataByTag %s", err.Error())
	}
	expected := uint(1)
	result, err := GetExifUIntDataByTag(x, exif.Orientation)
	if err != nil {
		t.Errorf("Error on TestGetExifUIntDataByTag %s", err.Error())
	}
	assert.NotNil(result, t)
	assert.Equals(expected, result, t)
}

func TestGetExifFloatDataByTag(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/images/DSC03783.JPG")
	x, err := GetExifFromBytes(file)
	if err != nil {
		t.Errorf("Error on GetExifFloatDataByTag %s", err.Error())
	}
	expected := float64(14.2)
	result, err := GetExifFloatDataByTag(x, exif.FocalLength)
	if err != nil {
		t.Errorf("Error on GetExifFloatDataByTag %s", err.Error())
	}
	assert.NotNil(result, t)
	assert.Equals(expected, result, t)
}
func TestGetStringExposureMode(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/images/DSC03783.JPG")
	x, err := GetExifFromBytes(file)
	if err != nil {
		t.Errorf("Error on GetStringExposureMode %s", err.Error())
	}
	expected := "manual"
	result, err := GetStringExposureMode(x)
	if err != nil {
		t.Errorf("Error on GetStringExposureMode %s", err.Error())
	}
	assert.NotNil(result, t)
	assert.Equals(expected, result, t)
}
