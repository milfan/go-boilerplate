package api_error

import (
	"net/http"

	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
)

const (
	// this section is used for error on service infrastructure
	DEVICE_ID_REQUIRED       = "SERVICE_NAME-SRV-30001"
	FAILED_GET_CONTEXT_VALUE = "SERVICE_NAME-SRV-30002"
	FAILED_DECODE_JSON       = "SERVICE_NAME-SRV-30003"
)

func InfraErrorDicts() map[string]*pkg_errors.Error {
	return map[string]*pkg_errors.Error{
		DEVICE_ID_REQUIRED: {
			ClientMessage: "Device ID required",
			ErrorCode:     DEVICE_ID_REQUIRED,
			HttpCode:      http.StatusBadRequest,
		},
		FAILED_GET_CONTEXT_VALUE: {
			ClientMessage: "Something wrong when get context data",
			ErrorCode:     FAILED_GET_CONTEXT_VALUE,
			HttpCode:      http.StatusUnprocessableEntity,
		},
		FAILED_DECODE_JSON: {
			ClientMessage: "Failed when decode",
			ErrorCode:     FAILED_DECODE_JSON,
			HttpCode:      http.StatusUnprocessableEntity,
		},
	}
}
