package middleware

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/milfan/go-boilerplate/configs/constants"
	api_error "github.com/milfan/go-boilerplate/internal/api/errors"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
	"github.com/sirupsen/logrus"
)

func GatherRequestData(
	response pkg_response.IResponse,
	logger *logrus.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		contentType := ctx.GetHeader(constants.HTTP_CONTENT_TYPE)

		var requestData interface{}
		if contentType != constants.HTTP_APPLICATION_FORM {

			body, _ := io.ReadAll(ctx.Request.Body)
			if len(body) > 0 {
				if err := json.Unmarshal(body, &requestData); err != nil {
					commErr := pkg_errors.New().Error(api_error.INVALID_PAYLOAD_REQUEST, err)
					response.HttpError(ctx, commErr)
					return
				}
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}

		}

		headers := make(map[string]string)
		for key, values := range ctx.Request.Header {
			headers[key] = values[0] // Log only the first value of each header
		}

		ctx.Next()

		logger.WithFields(map[string]interface{}{
			"request_id":   uuid.NewString(),
			"method":       ctx.Request.Method,
			"path":         ctx.Request.URL.Path,
			"headers":      headers,
			"status":       ctx.Writer.Status(),
			"client_ip":    ctx.ClientIP(),
			"user_agent":   ctx.Request.UserAgent(),
			"request_uri":  ctx.Request.RequestURI,
			"request_body": requestData,
		}).Info(ctx.Request.URL.Path)
	}
}
