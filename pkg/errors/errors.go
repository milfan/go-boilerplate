package pkg_errors

import (
	"fmt"
	"net/http"
	"strings"

	ozzo_validation "github.com/go-ozzo/ozzo-validation/v4"
)

const UNKNOWN_ERROR = "-"

type ErrorRequest map[string]string

type (
	Error struct {
		ErrorCode     string       `json:"code"`
		HttpCode      int          `json:"httpCode"`
		ClientMessage string       `json:"message"`
		ErrTrace      error        `json:"-"`
		ErrRequest    ErrorRequest `json:"errorRequest,omitempty"`
		Meta          *Meta        `json:"meta,omitempty"`
	}
	Meta struct {
		Timestamp string `json:"timestamp"`
		RequestId string `json:"requestId"`
	}
)

func (c Error) Error() string {
	return fmt.Sprintf("CommonError: %s.", c.ClientMessage)
}

type (
	IErrors interface {
		Error(errCode string, traceErrorMessage error) *Error
		ErrorValidate(errCode string, errMessage interface{}) *Error
	}
	pkgError struct{}
)

// Error implements IErrors.
func (g *pkgError) Error(commonErrCode string, traceErrorMessage error) *Error {

	dictionaries := errorDicts.Errors[commonErrCode]

	// to check if error in not listed
	if dictionaries == nil {
		return &Error{
			ClientMessage: "Unknown error", // this unknown client message if the error not registered
			ErrorCode:     commonErrCode,
			HttpCode:      http.StatusInternalServerError,
			ErrTrace:      traceErrorMessage,
		}
	}

	return &Error{
		ClientMessage: dictionaries.ClientMessage,
		ErrorCode:     dictionaries.ErrorCode,
		ErrTrace:      traceErrorMessage,
		HttpCode:      dictionaries.HttpCode,
	}
}

// AddValidationError implements IErrors.
func (g *pkgError) ErrorValidate(
	commonErrCode string,
	errMessage interface{},
) *Error {
	dictionaries := errorDicts.Errors[commonErrCode]

	comErr := Error{
		ClientMessage: dictionaries.ClientMessage,
		ErrorCode:     dictionaries.ErrorCode,
		HttpCode:      dictionaries.HttpCode,
	}
	// to check if error in not listed
	if dictionaries == nil {
		return &Error{
			ClientMessage: "Unknown error", // this unknown client message if the error not registered
			ErrorCode:     commonErrCode,
			HttpCode:      http.StatusInternalServerError,
		}
	}
	if _err, ok := errMessage.(ozzo_validation.Errors); ok {
		comErr.ErrRequest = buildValidationError(_err)
	}

	return &comErr
}

func buildValidationError(err error) ErrorRequest {
	var errors ErrorRequest = map[string]string{}

	errValidate := strings.Split(err.Error(), ";")
	for _, err := range errValidate {
		errPerField := strings.Split(err, ":")
		if len(errPerField[0]) <= 1 {
			errors["error"] = errPerField[0]
		} else {
			errors[strings.TrimSpace(errPerField[0])] = strings.TrimSpace(errPerField[1])
		}
	}

	return errors
}

func New() IErrors {
	return &pkgError{}
}
