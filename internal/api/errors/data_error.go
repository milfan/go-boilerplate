package api_error

import (
	"net/http"

	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
)

const (
	// this section is used when database communication got error
	REPOSITORY_GET_ERROR    = "SERVICE_NAME-SRV-20001"
	REPOSITORY_SAVE_ERROR   = "SERVICE_NAME-SRV-20002"
	REPOSITORY_UPDATE_ERROR = "SERVICE_NAME-SRV-20003"
	ERROR_DUPLICATE_VALUE   = "SERVICE_NAME-SRV-20004"

	// this section is used to error collection/data empty
	DATA_EMPLOYEE_EMPTY = "SERVICE_NAME-SRV-22001"
)

func DataErrorDicts() map[string]*pkg_errors.Error {
	return map[string]*pkg_errors.Error{

		// this section is used when database communication got error
		REPOSITORY_GET_ERROR: {
			ClientMessage: "Something wrong when get data",
			ErrorCode:     REPOSITORY_GET_ERROR,
			HttpCode:      http.StatusUnprocessableEntity,
		},
		REPOSITORY_SAVE_ERROR: {
			ClientMessage: "Something wrong when save data",
			ErrorCode:     REPOSITORY_SAVE_ERROR,
			HttpCode:      http.StatusUnprocessableEntity,
		},
		REPOSITORY_UPDATE_ERROR: {
			ClientMessage: "Something wrong when update data",
			ErrorCode:     REPOSITORY_UPDATE_ERROR,
			HttpCode:      http.StatusUnprocessableEntity,
		},
		ERROR_DUPLICATE_VALUE: {
			ClientMessage: "Data is duplate",
			ErrorCode:     ERROR_DUPLICATE_VALUE,
			HttpCode:      http.StatusConflict,
		},

		// this section is used to error collection/data empty
		DATA_EMPLOYEE_EMPTY: {
			ClientMessage: "Employee data not found",
			ErrorCode:     DATA_EMPLOYEE_EMPTY,
			HttpCode:      http.StatusNotFound,
		},
	}
}
