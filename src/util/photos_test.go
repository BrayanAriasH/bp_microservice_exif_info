package util

import (
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
)

func TestSeparateStringGPSCoordinateOk(t *testing.T) {
	coordinate := "[\"73/1\",\"50/1\",\"4492/100\"]"
	expectedDegree := float64(73 / 1)
	expectedMinute := float64(50 / 1)
	expecteSecond := float64(4492 / 100)

	resultDegree, resultMinute, resultSecond, err := SeparateStringGPSCoordinate(coordinate)
	assert.Equals(err, nil, t)
	assert.Equals(expectedDegree, resultDegree, t)
	assert.Equals(expectedMinute, resultMinute, t)
	assert.Equals(expecteSecond, resultSecond, t)
}

func TestSeparateStringGPSCoordinateError1(t *testing.T) {
	coordinate := "[\"73/1\",\"50/1\",\"4492/100\", buuuuuuu ]"

	resultDegree, resultMinute, resultSecond, err := SeparateStringGPSCoordinate(coordinate)
	assert.Equals(float64(0), resultDegree, t)
	assert.Equals(float64(0), resultMinute, t)
	assert.Equals(float64(0), resultSecond, t)
	assert.NotNil(err, t)
	errExpected := NewError(constant.GPSCoordinateError, coordinate)
	assert.Equals(err.Error(), errExpected.Error(), t)
}

func TestSeparateStringGPSCoordinateError2(t *testing.T) {
	coordinate := string("[\"73/1\",\"50/1\",\"4492*100\"]")

	resultDegree, resultMinute, resultSecond, err := SeparateStringGPSCoordinate(coordinate)
	assert.Equals(float64(0), resultDegree, t)
	assert.Equals(float64(0), resultMinute, t)
	assert.Equals(float64(0), resultSecond, t)
	assert.NotNil(err, t)
	errExpected := NewError(constant.GPSCoordinateValueError, "\"4492*100\"]", coordinate)
	assert.Equals(errExpected.Error(), err.Error(), t)
}

func TestGetDecimalDegreeCoordinateOk(t *testing.T) {
	gpsLongitude := "[\"73/1\",\"50/1\",\"4492/100\"]"
	gpsLatitude := "[\"5/1\",\"25/1\",\"3256/100\"]"
	gpsLongitudeRef := "W"
	gpsLatitudeRef := "N"
	expectedLongitude := -73.84555555555555
	expectedLatitude := 5.4255555555555555
	resultLatitude, resultLongitude, err := GetDecimalDegreeCoordinate(gpsLatitude, gpsLongitude, gpsLatitudeRef, gpsLongitudeRef)
	assert.Nil(err, t)
	assert.Equals(expectedLatitude, resultLatitude, t)
	assert.Equals(expectedLongitude, resultLongitude, t)
}

func TestGetDecimalDegreeCoordinateError1(t *testing.T) {
	gpsLongitude := "[\"73/1\",\"50/1\",\"4492/100\"]"
	gpsLatitude := "[\"5/1\",\"25/1\",\"3256/100\"]"
	gpsLongitudeRef := "W"
	gpsLatitudeRef := "N"
	expectedLongitude := -73.84556
	expectedLatitude := 5.42556
	resultLatitude, resultLongitude, err := GetDecimalDegreeCoordinate(gpsLatitude, gpsLongitude, gpsLatitudeRef, gpsLongitudeRef)
	assert.Nil(err, t)
	assert.NotEquals(expectedLatitude, resultLatitude, t)
	assert.NotEquals(expectedLongitude, resultLongitude, t)
}
