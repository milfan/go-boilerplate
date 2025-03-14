package api_error

import (
	"net/http"

	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
)

const (
	// this section is used for error on handler api request
	REQUEST_TIME_OUT                = "SERVICE_NAME-SRV-40001"
	ERROR_VALIDATE_ENUM_VALUE       = "SERVICE_NAME-SRV-40002"
	ERROR_VALIDATE_PARSE_DATE_VALUE = "SERVICE_NAME-SRV-40003"
	ERROR_VALIDATE_PARSE_VALUE      = "SERVICE_NAME-SRV-40004"
	INVALID_PAYLOAD_REQUEST         = "SERVICE_NAME-SRV-40005"
	INVALID_DECODE_PAYLOAD_REQUEST  = "SERVICE_NAME-SRV-40006"
	MISSING_DEVICE_ID               = "SERVICE_NAME-SRV-40007"
	DEVICE_ID_TOO_LONG              = "SERVICE_NAME-SRV-40008"
)

func InterfaceErrorDicts() map[string]*pkg_errors.Error {
	return map[string]*pkg_errors.Error{
		REQUEST_TIME_OUT: {
			ClientMessage: "Request timeout reach limit",
			ErrorCode:     REQUEST_TIME_OUT,
			HttpCode:      http.StatusRequestTimeout,
		},
		ERROR_VALIDATE_ENUM_VALUE: {
			ClientMessage: "Invalid value",
			ErrorCode:     ERROR_VALIDATE_ENUM_VALUE,
			HttpCode:      http.StatusRequestTimeout,
		},
		ERROR_VALIDATE_PARSE_DATE_VALUE: {
			ClientMessage: "Invalid value",
			ErrorCode:     ERROR_VALIDATE_PARSE_DATE_VALUE,
			HttpCode:      http.StatusRequestTimeout,
		},
		ERROR_VALIDATE_PARSE_VALUE: {
			ClientMessage: "Invalid value",
			ErrorCode:     ERROR_VALIDATE_PARSE_VALUE,
			HttpCode:      http.StatusRequestTimeout,
		},
		INVALID_PAYLOAD_REQUEST: {
			ClientMessage: "Data yang dikirim tidak sesuai",
			ErrorCode:     INVALID_PAYLOAD_REQUEST,
			HttpCode:      http.StatusBadRequest,
		},
		INVALID_DECODE_PAYLOAD_REQUEST: {
			ClientMessage: "Gagal decode pada data yang dikirim",
			ErrorCode:     INVALID_DECODE_PAYLOAD_REQUEST,
			HttpCode:      http.StatusBadRequest,
		},

		MISSING_DEVICE_ID: {
			ClientMessage: "X-Device-ID is required",
			ErrorCode:     MISSING_DEVICE_ID,
			HttpCode:      http.StatusForbidden,
		},
		DEVICE_ID_TOO_LONG: {
			ClientMessage: "X-Device-ID is too long",
			ErrorCode:     DEVICE_ID_TOO_LONG,
			HttpCode:      http.StatusBadRequest,
		},
	}
}
