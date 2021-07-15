package model

import "encoding/json"

type Photo struct {
	MakeWith         string  `json:"make_with"`
	DateTimeOriginal string  `json:"date_time_original"`
	EpochDate        int64   `json:"epoch_date"`
	ExposureTime     float64 `json:"exposure_time"`
	Orientation      uint    `json:"orientation"`
	Model            string  `json:"model"`
	ExposureMode     string  `json:"exposure_mode"`
	FNumber          float64 `json:"f_number"`
	GPSLongitude     string  `json:"gps_longitude"`
	GPSLongitudeRef  string  `json:"gps_longitude_ref"`
	GPSLatitude      string  `json:"gps_latitude"`
	GPSLatitudeRef   string  `json:"gps_latitude_ref"`
	GPSAltitude      float64 `json:"gps_altitude"`
	PixelXDimension  uint    `json:"pixel_x_dimension"`
	PixelYDimension  uint    `json:"pixel_y_dimension"`
	DDLongitude      float64 `json:"dd_longitude"`
	DDLatitude       float64 `json:"dd_latitude"`
}

func (Photo) NewPhoto() *Photo {
	return &Photo{}
}

func (p *Photo) String() string {
	result, _ := json.Marshal(p)
	return string(result)
}

func (p *Photo) Validate() (err error) {
	if p.GPSLongitude != "" && p.GPSLatitude != "" {
		//Validate coordinates
		return nil
	}
	return nil
}
