package api_error

import (
	"net/http"

	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
)

const (
	// this section is used for logical error or business error
	SOME_ERROR = "SERVICE_NAME-SRV-10001"
)

func AppErrorDicts() map[string]*pkg_errors.Error {
	return map[string]*pkg_errors.Error{
		SOME_ERROR: {
			ClientMessage: "some error",
			ErrorCode:     SOME_ERROR,
			HttpCode:      http.StatusBadRequest,
		},
	}
}
