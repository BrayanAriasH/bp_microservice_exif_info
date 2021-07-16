package constant

type ErrorText string

const (
	FileMustBeNotNull       ErrorText = "The filename must be not null."
	FileNotPermitted        ErrorText = "The filename is not permitted."
	GPSCoordinateError      ErrorText = "The coordinate %s is not Ok."
	GPSCoordinateValueError ErrorText = "The value %s in coordinate %s is not expected."
	ImageFileMustNotBeNull  ErrorText = "The file image must be not null."
)
