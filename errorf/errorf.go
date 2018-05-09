// Package errors implements functions to manipulate errors.
// 完全参考官方errors的实现，增加格式化的功能
package errorf

import (
	"fmt"
)

// New returns an error that formats as the given text.
func New(messages ...interface{}) error {
	return &errorString{fmt.Sprint(messages...)}
}

func Newf(format string, messages ...interface{}) error {
	return &errorString{fmt.Sprintf(format, messages...)}
}

// errorString is a trivial implementation of error.
type errorString struct {
	msg string
}

func (e *errorString) Error() string {
	return e.msg
}
