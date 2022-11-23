package api

import "go.uber.org/zap"

var logger = zap.NewExample()

type simpleError struct {
	e error
}

func (s simpleError) Error() string {
	return s.e.Error()
}

func plainError(err error) zap.Field {
	return zap.Error(simpleError{e: err})
}
