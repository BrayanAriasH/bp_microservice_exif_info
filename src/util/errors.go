package util

import (
	"errors"
	"fmt"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
)

func NewError(errorDesc constant.ErrorText, extraInfo ...interface{}) (err error) {
	var finalString string
	if extraInfo != nil {
		finalString = fmt.Sprintf(string(errorDesc), extraInfo...)

	} else {
		finalString = string(errorDesc)
	}
	return errors.New(finalString)
}
