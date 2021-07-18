package model

import (
	"io/ioutil"
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
)

func TestCreatePhotoFromFileOk(t *testing.T) {
	file, err := ioutil.ReadFile("../../test/images/DSC03783.JPG")
	if err != nil {
		t.Errorf("Error reading file %v", err)
	}
	photo, err := CreatePhotoFromFile(file)
	if err != nil {
		t.Errorf("Error reading file %v", err)
	}
	assert.NotNil(photo, t)
	assert.Equals("vertical", photo.Orientation, t)
}
