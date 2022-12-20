package merror

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

type Bizerr struct {
	error error      // Wrapped error.
	text  string     // Custom Error text when Error is created, might be empty when its code is not nil.
	code  gcode.Code // Error code if necessary.
}

func New(text string) error {
	return &Bizerr{
		text: text,
		code: gcode.CodeNil,
	}
}

func (err *Bizerr) Error() string {
	if err == nil {
		return ""
	}
	errStr := err.text
	if errStr == "" && err.code != nil {
		errStr = err.code.Message()
	}
	if err.error != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += err.error.Error()
	}
	return errStr
}
