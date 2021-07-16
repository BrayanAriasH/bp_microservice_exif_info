package util

import (
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
)

func TestNewError(t *testing.T) {
	error := NewError(constant.GPSCoordinateValueError, "value1", "value2")
	assert.NotNil(error, t)
	expected := "The value value1 in coordinate value2 is not expected."
	assert.Equals(expected, error.Error(), t)
}
