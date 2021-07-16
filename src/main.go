package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/model"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type Printer struct{}

func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%40s: %s\n", name, tag)
	return nil
}
func main() {
	byteData, err := ioutil.ReadFile("./test/images/DSC04249.JPG")
	if err != nil {
		log.Panic(err)
	}
	photo, err := model.CreatePhotoFromFile(byteData)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Final photo:", photo)
}
