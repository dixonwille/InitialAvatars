package avatar

import "errors"

type avError struct {
	err   error
	value string
}

var (
	errValueOutOfRange = errors.New("Value is out of range.")
)

func (err avError) Error() string {
	if err.value != "" {
		return err.err.Error() + ": " + err.value
	}
	return err.err.Error()
}

func newError(err error, value string) *avError {
	return &avError{
		err:   err,
		value: value,
	}
}

//IsValueOutOfRange returns whether the error is a value out of range error.
func IsValueOutOfRange(err error) bool {
	e, ok := err.(avError)
	if !ok {
		return false
	}
	if e.err.Error() != errValueOutOfRange.Error() {
		return false
	}
	return true
}
