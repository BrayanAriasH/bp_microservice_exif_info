package util

import (
	"io/ioutil"
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
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
