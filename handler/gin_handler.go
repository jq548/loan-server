package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"loan-server/common/errors"
	"loan-server/common/res"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// GinLogger Receive the default log of the gin framework
func GinLogger() gin.HandlerFunc {
	logger := zap.L()
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery gin exception handling
func GinRecovery(stack bool) gin.HandlerFunc {
	logger := zap.L()
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					failed := res.Failed(errors.NetworkAnomaly)
					c.JSON(failed.Code, failed)
					//c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}

				switch err.(type) {
				case *errors.Err:
					e := err.(*errors.Err)
					failed := res.Failed(e.Code)
					c.JSON(failed.Code, failed)
					c.Abort()
				default:
					unknownErr := res.UnknownErr(nil)
					c.JSON(http.StatusOK, unknownErr)
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}

// HandleNotFound 404
func HandleNotFound(c *gin.Context) {
	zap.S().Errorf("handle not found: %v", c.Request.RequestURI)
	c.JSON(http.StatusNotFound, res.Failed(errors.UriNotFoundOrMethodNotSupport))
	return
}
