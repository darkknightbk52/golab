package error_handling

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorType uint

const (
	NoType ErrorType = iota
	BadRequest
	NotFound
)

type customError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
}

type errorContext struct {
	field   string
	message string
}

func (e customError) Error() string {
	return e.originalError.Error()
}

func (et ErrorType) New(msg string) error {
	return et.Newf(msg)
}

func (et ErrorType) Newf(msg string, args ...interface{}) error {
	return &customError{
		errorType:     et,
		originalError: fmt.Errorf(msg, args...),
	}
}

func (et ErrorType) Wrap(err error, msg string) error {
	return et.Wrapf(err, msg)
}

func (et ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return &customError{
		errorType:     et,
		originalError: errors.Wrapf(err, msg, args...),
	}
}

func New(msg string) error {
	return Newf(msg)
}

func Newf(msg string, args ...interface{}) error {
	return &customError{
		errorType:     NoType,
		originalError: fmt.Errorf(msg, args...),
	}
}

func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

func Wrapf(err error, msg string, args ...interface{}) error {
	customErr, ok := err.(*customError)
	if !ok {
		return &customError{
			errorType:     NoType,
			originalError: errors.Wrapf(err, msg, args...),
		}
	}

	return &customError{
		errorType:     customErr.errorType,
		originalError: err,
		context:       customErr.context,
	}
}

func Cause(err error) string {
	return errors.Cause(err).Error()
}

func AddErrorContext(err error, field, msg string) error {
	errCtx := errorContext{field: field, message: msg}
	wrappedErr, ok := err.(*customError)
	if !ok {
		return &customError{
			errorType:     NoType,
			originalError: err,
			context:       errCtx,
		}
	}

	return &customError{
		errorType:     wrappedErr.errorType,
		originalError: wrappedErr,
		context:       errCtx,
	}
}

func GetErrorContext(err error) map[string]string {
	wrappedErr, ok := err.(*customError)
	if !ok {
		return nil
	}

	return map[string]string{"field": wrappedErr.context.field, "message": wrappedErr.context.message}
}

func GetType(err error) ErrorType {
	wrappedErr, ok := err.(*customError)
	if !ok {
		return NoType
	}
	return wrappedErr.errorType
}
