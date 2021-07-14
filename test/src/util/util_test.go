package util_test

import (
	"io/ioutil"
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
)

func TestValidateFileType(t *testing.T) {
	file, _ := ioutil.ReadFile("../../images/DSC03986-min.jpg")
	isValid, err := util.ValidateFileType(file)
	if err != nil {
		t.Errorf("Error in TestValidateFileType: %v", err)
	}
	if !isValid {
		t.Errorf("Expected value: %v, actual value %v", true, isValid)
	}
}
