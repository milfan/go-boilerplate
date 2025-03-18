package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	api_error "github.com/milfan/go-boilerplate/internal/api/errors"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
)

func RequestTimeoutMiddleware(
	timeout time.Duration,
	response pkg_response.IResponse,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace the original request's context with the new context
		c.Request = c.Request.WithContext(ctx)

		// Create a channel to signal when the request is done
		done := make(chan struct{})

		go func() {
			c.Next() // Continue with the next middleware or the handler
			close(done)
		}()

		// Wait for either the request to finish or the context to timeout
		select {
		case <-done:
			// Request completed within the timeout
			return
		case <-ctx.Done():
			// Timeout reached, abort the request
			commErr := pkg_errors.New().Error(api_error.REQUEST_TIME_OUT, nil)
			response.HttpError(c, commErr)
			return
		}
	}
}
