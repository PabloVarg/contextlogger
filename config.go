package contextlogger

import (
	"log/slog"
	"net/http"
)

type (
	BucketConfigurator = func(c *BucketConfig)
	BucketConfig       struct {
		capacity int
	}
)

func defaultBucketConfig() *BucketConfig {
	return &BucketConfig{
		capacity: 20,
	}
}

func WithCapacity(capacity int) BucketConfigurator {
	return func(c *BucketConfig) {
		c.capacity = capacity
	}
}

type (
	middlewareConfigurator = func(c *middlewareConfig)
	middlewareConfig       struct {
		logger            *slog.Logger
		withDefaultValues bool
		preHook           func(w http.ResponseWriter, r *http.Request)
		postHook          func(w http.ResponseWriter, r *http.Request)
	}
)

func defaultMiddlewareConfig() *middlewareConfig {
	return &middlewareConfig{
		logger:            slog.Default(),
		withDefaultValues: true,
		preHook:           func(w http.ResponseWriter, r *http.Request) {},
		postHook:          func(w http.ResponseWriter, r *http.Request) {},
	}
}

func WithLogger(logger *slog.Logger) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.logger = logger
	}
}

func WithDefaultValues(enable bool) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.withDefaultValues = enable
	}
}

func WithPreHook(hook func(w http.ResponseWriter, r *http.Request)) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.preHook = hook
	}
}

func WithPostHook(hook func(w http.ResponseWriter, r *http.Request)) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.postHook = hook
	}
}
