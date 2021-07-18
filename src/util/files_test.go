package util_test

import (
	"io/ioutil"
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
)

func TestValidateFileTypeNoError(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/images/DSC04249.JPG")
	isValid, err := util.ValidateFileType(file)
	if err != nil {
		t.Errorf("Error in TestValidateFileType: %v", err)
	}
	if !isValid {
		t.Errorf("Expected value: %v, actual value %v", true, isValid)
	}
}

func TestValidateFileTypeNil(t *testing.T) {
	isValid, err := util.ValidateFileType(nil)
	if isValid {
		t.Errorf("Expected isValid value: %v, actual value %v", false, isValid)
	}
	if err == nil {
		t.Errorf("Expected err value: %v, actual value %v", nil, err)
	}
}
