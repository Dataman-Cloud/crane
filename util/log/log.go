package log

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

var (
	// G is an alias for GetLogger.
	//
	// We may want to define this locally to a package to get package tagged log
	// messages.
	G = GetLogger

	// L is an alias for the the standard logger.
	L = logrus.NewEntry(logrus.StandardLogger())
)

type loggerKey struct{}

// WithLogger returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger retrieves the current logger from the context. If no logger is
// available, the default logger is returned.
func GetLogger(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})

	if logger == nil {
		return L
	}

	return logger.(*logrus.Entry)
}

func Ginrus(logger *logrus.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:      true,
			DisableColors:    false,
			DisableTimestamp: false,
			FullTimestamp:    true,
			DisableSorting:   false,
		})

		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		entry := logrus.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(timeFormat),
		})
		_ = entry

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			//entry.Error(c.Errors.String())
			logger.Errorf("%d %s %s %s %s %s %s %s",
				c.Writer.Status(),
				c.Request.Method,
				path,
				c.ClientIP(),
				latency,
				c.Request.UserAgent(),
				end.Format(timeFormat),
				c.Errors.String(),
			)
		} else {
			//entry.Info()
			logger.Infof("%d %s %s %s %s %s %s",
				c.Writer.Status(),
				c.Request.Method,
				path,
				c.ClientIP(),
				latency,
				c.Request.UserAgent(),
				end.Format(timeFormat),
			)
		}
	}
}
