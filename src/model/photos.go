package model

import (
	"encoding/json"
	"time"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
	"github.com/google/uuid"
	"github.com/rwcarlsen/goexif/exif"
)

type Photo struct {
	Id               string    `json:"id"`
	MakeWith         string    `json:"make_with"`
	DateTimeOriginal time.Time `json:"date_time_original"`
	ExposureTime     string    `json:"exposure_time"`
	Orientation      string    `json:"orientation"`
	Model            string    `json:"model"`
	ExposureMode     string    `json:"exposure_mode"`
	FNumber          float64   `json:"f_number"`
	Longitude        float64   `json:"longitude"`
	Latitude         float64   `json:"latitude"`
	PixelXDimension  uint      `json:"resolution_x"`
	PixelYDimension  uint      `json:"resolution_y"`
	ISO              uint      `json:"iso"`
}

func NewPhoto() *Photo {
	return &Photo{}
}

func (p *Photo) String() string {
	result, _ := json.Marshal(p)
	return string(result)
}

func CreatePhotoFromFile(data []byte) (photo *Photo, err error) {
	if data == nil {
		return nil, util.NewError(constant.FileMustBeNotNull)
	}
	x, err := util.GetExifFromBytes(data)
	if err != nil {
		return nil, err
	}
	photo = NewPhoto()
	photo.MakeWith, err = util.GetExifStringDataByTag(x, exif.Make)
	if err != nil {
		return nil, err
	}
	photo.DateTimeOriginal, err = x.DateTime()
	if err != nil {
		return nil, err
	}
	photo.ExposureTime, err = util.GetExifStringDataByTag(x, exif.ExposureTime)
	if err != nil {
		return nil, err
	}
	photo.Orientation, err = util.GetStringOrientation(x)
	if err != nil {
		return nil, err
	}
	photo.Model, err = util.GetExifStringDataByTag(x, exif.Model)
	if err != nil {
		return nil, err
	}
	photo.ExposureMode, err = util.GetStringExposureMode(x)
	if err != nil {
		return nil, err
	}
	photo.FNumber, err = util.GetExifFloatDataByTag(x, exif.FNumber)
	if err != nil {
		return nil, err
	}
	photo.Latitude, photo.Longitude, err = x.LatLong()
	if err != nil {
		return nil, err
	}
	photo.PixelXDimension, err = util.GetExifUIntDataByTag(x, exif.PixelXDimension)
	if err != nil {
		return nil, err
	}
	photo.PixelYDimension, err = util.GetExifUIntDataByTag(x, exif.PixelYDimension)
	if err != nil {
		return nil, err
	}
	photo.ISO, err = util.GetExifUIntDataByTag(x, exif.ISOSpeedRatings)
	if err != nil {
		return nil, err
	}
	photo.Id = uuid.New().String()
	return photo, nil
}
