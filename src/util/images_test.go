package util

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/assert"
)

func TestCompressImage(t *testing.T) {
	paths := []string{"../../test/images/DSC03783.JPG"}
	for _, path := range paths {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("Error on TestCompressImage: %s", err)
			return
		}
		img, err := CompressImage(file, 800, 0)
		assert.Nil(err, t)
		assert.NotNil(img, t)
	}
}
