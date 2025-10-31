package httpmiddleware

import (
	"log/slog"
	"net/http"
)

type (
	// middlewareConfigurator mutates the configuration to create a
	// middleware with contextual logging
	middlewareConfigurator = func(c *middlewareConfig)
	// middlewareConfig stores the configuration for a middleware
	middlewareConfig struct {
		// Instance of slog.Logger to use for writting logs
		logger *slog.Logger
		// Message to use for logging, it's directly passed as the message for
		// the Log method for a slog.Logger
		message string
		// Indicates whether to include some basic values without the
		// user having to define them. This will include http method, path, etc.
		withDefaultValues bool
		// preHook is a function that executes before the HTTP request is processed.
		// It receives the http.Request and can be used to perform setup operations
		// or add contextual information before logging.
		preHook func(r *http.Request)
		// postHook is a function that executes after the HTTP request is processed.
		// It receives the http.Request and can be used to perform cleanup operations
		// or add additional contextual information after processing the request.
		postHook func(r *http.Request)
	}
)

// DefaultMiddlewareConfig returns a middlewareConfig with sensible default values.
func DefaultMiddlewareConfig() *middlewareConfig {
	return &middlewareConfig{
		logger:            slog.Default(),
		message:           "http server hit",
		withDefaultValues: true,
		preHook:           func(r *http.Request) {},
		postHook:          func(r *http.Request) {},
	}
}

// WithLogger returns a middlewareConfigurator that sets a custom slog.Logger
// for the middleware. This allows you to use a specific logger instance instead
// of the default logger.
func WithLogger(logger *slog.Logger) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.logger = logger
	}
}

// WithMessage returns a middlewareConfigurator that sets a custom log message
// for the middleware. This message will be used as the primary log message when
// the middleware logs HTTP requests.
func WithMessage(message string) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.message = message
	}
}

// WithDefaultValues returns a middlewareConfigurator that enables or disables
// the automatic inclusion of default HTTP request values in logs. When enabled,
// the middleware will automatically log standard HTTP information such as method,
// path, and other request details.
func WithDefaultValues(enable bool) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.withDefaultValues = enable
	}
}

// WithPreHook returns a middlewareConfigurator that sets a function to be executed
// before the HTTP request is processed. The hook function receives the http.Request
// and can be used to add custom logic, modify context, or perform setup operations
// before logging occurs.
func WithPreHook(hook func(r *http.Request)) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.preHook = hook
	}
}

// WithPostHook returns a middlewareConfigurator that sets a function to be executed
// after the HTTP request is processed. The hook function receives the http.Request
// and can be used to add custom logic, perform cleanup operations, or add additional
// contextual information to logs after the request has been processed.
func WithPostHook(hook func(r *http.Request)) middlewareConfigurator {
	return func(c *middlewareConfig) {
		c.postHook = hook
	}
}
