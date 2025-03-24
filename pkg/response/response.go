package pkg_response

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pkg_constants "github.com/milfan/go-boilerplate/pkg/constants"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Meta struct {
	Page      int `json:"page,omitempty"`
	PerPage   int `json:"perPage,omitempty"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage,omitempty"`
}

type ResponseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type IResponse interface {
	// It used to return http response when it's ok
	HttpJSON(ctx *gin.Context, message string, data interface{}, meta *Meta)

	// It used to return http response when has error
	HttpError(ctx *gin.Context, err error)

	// It used to build meta response
	BuildMeta(page int, perPage int, count int64) *Meta
}

type response struct {
	logger *logrus.Logger
}

// BuildMeta implements IResponse.
func (r *response) BuildMeta(page int, perPage int, count int64) *Meta {
	x := math.Ceil(float64(count) / float64(perPage))
	totalPage := int(x)
	return &Meta{
		Page:      page,
		PerPage:   perPage,
		Total:     int(count),
		TotalPage: totalPage,
	}
}

// HttpError implements IResponse.
func (r *response) HttpError(ctx *gin.Context, err error) {

	var (
		requestID   string
		headers     interface{}
		requestData interface{}
	)

	// get request id from middleware
	getRequestID, _ := ctx.Get(pkg_constants.REQUEST_ID)
	if getRequestID != nil {
		requestID = getRequestID.(string)
	}

	// get request headers from middleware
	getHeaders, _ := ctx.Get(pkg_constants.REQUEST_HEADER)
	if getHeaders != nil {
		headers = getHeaders
	}

	// get request data from middleware
	getRequestData, _ := ctx.Get(pkg_constants.REQUEST_DATA)
	if getRequestData != nil {
		requestData = getRequestData
	}

	respError := pkg_errors.New().Error(pkg_errors.UNKNOWN_ERROR, nil)
	cerr, ok := err.(*pkg_errors.Error)
	if ok {
		respError = cerr
	}

	// add meta error
	respError.Meta = &pkg_errors.Meta{
		Timestamp: time.Now().Format(pkg_constants.SQLTimestampFormat),
		RequestId: requestID,
	}
	resp, _ := json.Marshal(respError)

	r.logger.WithFields(map[string]interface{}{
		"request_id":   requestID,
		"method":       ctx.Request.Method,
		"http_status":  respError.HttpCode,
		"headers":      headers,
		"client_ip":    ctx.ClientIP(),
		"user_agent":   ctx.Request.UserAgent(),
		"request_uri":  ctx.Request.RequestURI,
		"request_body": requestData,
		"error_code":   respError.ErrorCode,
		"trace_error":  respError.ErrTrace.Error(),
	}).Error(respError.ClientMessage)

	ctx.Writer.Header().Set(pkg_constants.CONTENT_TYPE, pkg_constants.CONTENT_TYPE_JSON)
	ctx.Writer.WriteHeader(respError.HttpCode)
	ctx.Writer.Write(resp)
	ctx.Abort()
}

// HttpJSON implements IResponse.
func (r *response) HttpJSON(ctx *gin.Context, message string, data interface{}, meta *Meta) {

	var (
		requestID   string
		headers     interface{}
		requestData interface{}
	)

	// get request id from middleware
	getRequestID, _ := ctx.Get(pkg_constants.REQUEST_ID)
	if getRequestID != nil {
		requestID = getRequestID.(string)
	}

	// get request headers from middleware
	getHeaders, _ := ctx.Get(pkg_constants.REQUEST_HEADER)
	if getHeaders != nil {
		headers = getHeaders
	}

	// get request data from middleware
	getRequestData, _ := ctx.Get(pkg_constants.REQUEST_DATA)
	if getRequestData != nil {
		requestData = getRequestData
	}

	response := ResponseMessage{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	resp, _ := json.Marshal(response)

	r.logger.WithFields(map[string]interface{}{
		"request_id":       requestID,
		"method":           ctx.Request.Method,
		"http_status":      http.StatusOK,
		"headers":          headers,
		"client_ip":        ctx.ClientIP(),
		"user_agent":       ctx.Request.UserAgent(),
		"request_uri":      ctx.Request.RequestURI,
		"request_body":     requestData,
		"request_response": response,
	}).Info(ctx.Request.URL.Path)

	ctx.Writer.Header().Set(pkg_constants.CONTENT_TYPE, pkg_constants.CONTENT_TYPE_JSON)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(resp)
}

func New(
	logger *logrus.Logger,
) IResponse {
	return &response{
		logger: logger,
	}
}
