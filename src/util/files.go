package util

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
)

//ValidateFileType check if a byte array is a image permitted type
func ValidateFileType(bytes []byte) (isValid bool, err error) {
	if bytes == nil {
		return false, NewError(constant.FileMustBeNotNull)
	}
	fileType := http.DetectContentType(bytes)
	log.Println(fmt.Sprintf("The file type is %s", fileType))
	if fileType != constant.JpgType {
		return false, NewError(constant.FileNotPermitted)
	}
	return true, nil
}
