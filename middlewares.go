package contextlogger

import (
	"log/slog"
	"net/http"
	"time"
)

func HttpMiddleware(next http.Handler, opts ...middlewareConfigurator) http.Handler {
	conf := defaultMiddlewareConfig()
	for _, opt := range opts {
		opt(conf)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		r = r.WithContext(EmbedLoggingAttrs(r.Context()))
		if conf.withDefaultValues {
			UpdateContext(
				r.Context(),
				"method", r.Method,
				"host", r.Host,
				"path", r.Pattern,
				"url", r.URL.String(),
			)
		}
		conf.preHook(w, r)

		defer func() {
			if conf.withDefaultValues {
				UpdateContext(
					r.Context(),
					"duration", time.Since(start).Nanoseconds(),
				)
			}
			conf.postHook(w, r)

			LogWithContext(r.Context(), conf.logger, slog.LevelInfo, "http server hit")
		}()

		next.ServeHTTP(w, r)
	})
}
