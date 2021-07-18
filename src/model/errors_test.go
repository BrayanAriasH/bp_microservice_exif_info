package model

import (
	"errors"
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
)

func TestNewErrorResponse(t *testing.T) {
	response := NewErrorResponse(1, errors.New("Hello error"))
	assert.NotNil(response, t)
	assert.Equals(uint(1), response.Code, t)
}
