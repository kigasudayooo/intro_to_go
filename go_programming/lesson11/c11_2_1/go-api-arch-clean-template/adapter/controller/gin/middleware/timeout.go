package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"

	"go-api-arch-clean-template/adapter/controller/gin/presenter"
)

func TimeoutMiddleware(duration time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(duration),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.JSON(presenter.NewErrorResponse(http.StatusRequestTimeout, "timeout"))
			c.Abort()
		}),
	)
}
