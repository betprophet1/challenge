package failure

import (
	"fmt"
)

type failure struct {
	err       error
	errorcode string
	code      int
}

func (f failure) Error() string {
	return fmt.Sprintf("status=%s | code=%d", f.errorcode, f.code)
}

func (f failure) Code() int {
	return f.code
}

func (f failure) ErrorCode() string {
	return f.errorcode
}

func (f failure) Err() error {
	return f.err
}

type Failure interface {
	Code() int
	ErrorCode() string
	Err() error
}

func NewFailure(errorcode string, code int) ErrorWraper {
	return func(err error) error {
		return failure{
			code:      code,
			errorcode: errorcode,
			err:       err,
		}
	}
}
