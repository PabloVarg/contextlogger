package contextlogger

import (
	"log/slog"
	"net/http"
)

type (
	// bucketConfigurator mutates the configuration to create a
	// new internal map where key / values will be stored.
	bucketConfigurator = func(c *bucketConfig)
	// bucketConfig stores the configuration options used to
	// create a new internal map where key / values will be stored.
	bucketConfig struct {
		// logger used to log with context
		logger *slog.Logger
		// Allocated capacity for a new bucket, this is sent to 'make' as
		// the capacity for internal maps. i.e. make(map[string]any, capacity)
		capacity int
	}
)

// Generate default values for a BucketConfig.
func defaultBucketConfig() *bucketConfig {
	return &bucketConfig{
		capacity: 20,
	}
}

// WithCapacity modifies the default capacity for a bucket.
// Returns a BucketConfigurator with the given capacity.
func WithCapacity(capacity int) bucketConfigurator {
	return func(c *bucketConfig) {
		c.capacity = capacity
	}
}

type (
	// middlewareConfigurator mutates the configuration to create a
	// middleware with contextual logging
	middlewareConfigurator = func(c *middlewareConfig)
	// middlewareConfig stores the configuration for a middleware
	middlewareConfig struct {
		// instance of slog.Logger to use for writting logs
		logger *slog.Logger
		// Indicates whether to include some basic values without the
		// user having to define them. This will include http method, path, etc.
		withDefaultValues bool
		//
		preHook  func(r *http.Request)
		postHook func(r *http.Request)
	}
)

func defaultMiddlewareConfig() *middlewareConfig {
	return &middlewareConfig{
		logger:            slog.Default(),
		withDefaultValues: true,
		preHook:           func(r *http.Request) {},
		postHook:          func(r *http.Request) {},
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

func WithPreHook(hook func(r *http.Request)) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.preHook = hook
	}
}

func WithPostHook(hook func(r *http.Request)) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.postHook = hook
	}
}
