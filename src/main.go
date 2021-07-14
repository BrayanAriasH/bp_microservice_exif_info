package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
)

func main() {
	fileByte, err := ioutil.ReadFile("./test/images/DSC03986-min.jpg")
	if err != nil {
		fmt.Printf("Error on Stats %v", err)
		return
	}
	isValid, err := util.ValidateFileType(fileByte)
	if err != nil {
		fmt.Printf("Error on ValidateFileType %v", err)
		return
	}
	fmt.Println("Is valid:", isValid)
}
