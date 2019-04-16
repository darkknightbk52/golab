package error_handling

import (
	"errors"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestErrorType_New(t *testing.T) {
	RegisterTestingT(t)

	testCases := []struct {
		errorType ErrorType
		err       *customError
		cause     string
	}{
		{
			errorType: NoType,
			err:       NoType.New("NoType error").(*customError),
			cause:     "NoType error",
		},
		{
			errorType: BadRequest,
			err:       BadRequest.New("BadRequest error").(*customError),
			cause:     "BadRequest error",
		},
		{
			errorType: NotFound,
			err:       NotFound.New("NotFound error").(*customError),
			cause:     "NotFound error",
		},
	}

	for _, c := range testCases {
		Expect(c.err.errorType).Should(Equal(c.errorType))
		Expect(Cause(c.err)).Should(Equal(c.cause))
	}
}

func TestErrorType_Newf(t *testing.T) {
	RegisterTestingT(t)

	testCases := []struct {
		errorType ErrorType
		err       *customError
		cause     string
	}{
		{
			errorType: NoType,
			err:       NoType.Newf("NoType %s", "error").(*customError),
			cause:     "NoType error",
		},
		{
			errorType: BadRequest,
			err:       BadRequest.Newf("BadRequest %s", "error").(*customError),
			cause:     "BadRequest error",
		},
		{
			errorType: NotFound,
			err:       NotFound.Newf("NotFound %s", "error").(*customError),
			cause:     "NotFound error",
		},
	}

	for _, c := range testCases {
		Expect(c.err.errorType).Should(Equal(c.errorType))
		Expect(Cause(c.err)).Should(Equal(c.cause))
	}
}

func TestErrorType_Wrap(t *testing.T) {
	RegisterTestingT(t)

	var osErr error
	osErr = &os.PathError{
		Op:   "remove",
		Path: "imagedDir",
		Err:  errors.New("no such file or directory"),
	}

	testCases := []struct {
		errorType ErrorType
		err       *customError
		cause     string
	}{
		{
			errorType: NoType,
			err:       NoType.Wrap(osErr, "NoType error").(*customError),
			cause:     "NoType error: remove imagedDir: no such file or directory",
		},
		{
			errorType: BadRequest,
			err:       BadRequest.Wrap(osErr, "BadRequest error").(*customError),
			cause:     "BadRequest error: remove imagedDir: no such file or directory",
		},
		{
			errorType: NotFound,
			err:       NotFound.Wrap(osErr, "NotFound error").(*customError),
			cause:     "NotFound error: remove imagedDir: no such file or directory",
		},
	}

	for _, c := range testCases {
		Expect(c.err.errorType).Should(Equal(c.errorType))
		Expect(Cause(c.err)).Should(Equal(c.cause))
	}
}

func TestErrorType_Wrapf(t *testing.T) {
	RegisterTestingT(t)

	var osErr error
	osErr = &os.PathError{
		Op:   "remove",
		Path: "imagedDir",
		Err:  errors.New("no such file or directory"),
	}

	testCases := []struct {
		errorType ErrorType
		err       *customError
		cause     string
	}{
		{
			errorType: NoType,
			err:       NoType.Wrapf(osErr, "NoType %s", "error").(*customError),
			cause:     "NoType error: remove imagedDir: no such file or directory",
		},
		{
			errorType: BadRequest,
			err:       BadRequest.Wrapf(osErr, "BadRequest %s", "error").(*customError),
			cause:     "BadRequest error: remove imagedDir: no such file or directory",
		},
		{
			errorType: NotFound,
			err:       NotFound.Wrapf(osErr, "NotFound %s", "error").(*customError),
			cause:     "NotFound error: remove imagedDir: no such file or directory",
		},
	}

	for _, c := range testCases {
		Expect(c.err.errorType).Should(Equal(c.errorType))
		Expect(Cause(c.err)).Should(Equal(c.cause))
	}
}
