package httpmiddleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/pablovarg/contextlogger"
	"github.com/pablovarg/contextlogger/buckets"
)

// HttpMiddleware is meant to be used with the package net/http from Go's standard library
func HttpMiddleware(next http.Handler, opts ...middlewareConfigurator) http.Handler {
	conf := DefaultMiddlewareConfig()
	for _, opt := range opts {
		opt(conf)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		r = r.WithContext(contextlogger.EmbedLogger(r.Context(), buckets.WithLogger(conf.logger)))
		if conf.withDefaultValues {
			contextlogger.UpdateContext(
				r.Context(),
				"method", r.Method,
				"host", r.Host,
				"path", r.Pattern,
				"url", r.URL.String(),
			)
		}
		conf.preHook(r)

		defer func() {
			if conf.withDefaultValues {
				contextlogger.UpdateContext(
					r.Context(),
					"duration", time.Since(start).Nanoseconds(),
				)
			}
			conf.postHook(r)

			contextlogger.LogWithContext(
				r.Context(),
				slog.LevelInfo,
				conf.message,
			)
		}()

		next.ServeHTTP(w, r)
	})
}
