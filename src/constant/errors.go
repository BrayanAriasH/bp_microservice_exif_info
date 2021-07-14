package constant

type ErrorText string

const (
	FileMustBeNotNull ErrorText = "The filename must be not null."
	FileNotPermitted  ErrorText = "The filename is not permitted."
)
