package errorx

import "reflect"

// Error error
type Error string

func (e Error) Error() string {
	return reflect.ValueOf(e).String()
}

const (
	ErrOriginFileFormatNotSupport Error = "origin file format is not supported"
	ErrTargetFileFormatNotSupport Error = "target file format is not supported"

	ErrOriginFileNotExist Error = "origin file is not exist"
	ErrTargetFileNotExist Error = "target file is not exist"

	ErrCovertFailed Error = "file conversion failed"
)
