package middleware

import (
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger() tele.MiddlewareFunc {
	appLogger := zap.L()

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			appLogger.Debug("telebot log middleware", zap.Any("context", c.Update()))
			return next(c)
		}
	}
}
