package model

import (
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
)

func TestNewRequestPhoto(t *testing.T) {
	request := NewRequestPhoto()
	assert.NotNil(request, t)
}
