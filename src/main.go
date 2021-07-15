package main

import (
	"fmt"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type Printer struct{}

func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%40s: %s\n", name, tag)
	return nil
}
func main() {
	// byteData, err := ioutil.ReadFile("./test/images/DSC04249.JPG")
	// if err != nil {
	// 	fmt.Printf("Error on ReadFile %v", err)
	// 	return
	// }
	// x, err := exif.Decode(bytes.NewReader(byteData))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var p Printer
	// x.Walk(p)
	data := "[\"73/1\",\"50/1\",\"4492/100\"]"
	splited := strings.Split(data, ",")
	for i := range splited {
		splited[i] = strings.ReplaceAll(splited[i], "\"", "")
		splited[i] = strings.ReplaceAll(splited[i], "[", "")
		splited[i] = strings.ReplaceAll(splited[i], "]", "")
		fmt.Println("Index:", i, " - Splited:", splited[i])
	}
}
