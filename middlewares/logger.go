package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s request \"%s\" |%s| -> %d, [%s].\n",
			params.TimeStamp.Format(time.RFC822Z),
			params.ClientIP,
			params.Path,
			params.Method,
			params.StatusCode,
			params.Latency,
		)
	})
}
