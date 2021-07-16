package model

type RequestPhoto struct {
	FileName      string `json:"file-name"`
	Base64Content string `json:"content"`
}

func NewRequestPhoto() (requestPhoto *RequestPhoto) {
	return &RequestPhoto{}
}
