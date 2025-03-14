package pkg_response

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pkg_constants "github.com/milfan/go-boilerplate/pkg/constants"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
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

type response struct{}

// BuildMeta implements IResponseClient.
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

// HttpError implements IResponseClient.
func (r *response) HttpError(ctx *gin.Context, err error) {

	// get request id from middleware
	getRequestID, _ := ctx.Get(pkg_constants.REQUEST_ID)
	requestID := ""
	if getRequestID != nil {
		requestID = getRequestID.(string)
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

	ctx.Writer.Header().Set(pkg_constants.CONTENT_TYPE, pkg_constants.CONTENT_TYPE_JSON)
	ctx.Writer.WriteHeader(respError.HttpCode)
	ctx.Writer.Write(resp)
	ctx.Abort()
}

// HttpJSON implements IResponseClient.
func (r *response) HttpJSON(ctx *gin.Context, message string, data interface{}, meta *Meta) {

	response := ResponseMessage{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	resp, _ := json.Marshal(response)

	ctx.Writer.Header().Set(pkg_constants.CONTENT_TYPE, pkg_constants.CONTENT_TYPE_JSON)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(resp)
}

func New() IResponse {
	return &response{}
}
