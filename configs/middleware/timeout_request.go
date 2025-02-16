package conf_middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestTimeoutMiddleware(
	timeout time.Duration,
	logger *logrus.Logger,
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
			c.JSON(http.StatusGatewayTimeout, "timeout")
			return
		}
	}
}
