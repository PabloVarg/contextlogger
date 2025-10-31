package buckets

import (
	"log/slog"
)

type (
	// BucketConfigurator mutates the configuration to create a
	// new internal map where key / values will be stored.
	BucketConfigurator = func(c *BucketConfig)
	// BucketConfig stores the configuration options used to
	// create a new internal map where key / values will be stored.
	BucketConfig struct {
		// Logger used to log with context
		Logger *slog.Logger
		// Allocated Capacity for a new bucket, this is sent to 'make' as
		// the Capacity for internal maps. i.e. make(map[string]any, Capacity)
		Capacity int
	}
)

// Generate default values for a BucketConfig.
func DefaultBucketConfig() *BucketConfig {
	return &BucketConfig{
		Capacity: 20,
	}
}

// WithLogger modifies the default logger for a bucket.
// Returns a BucketConfigurator with the given logger.
func WithLogger(logger *slog.Logger) BucketConfigurator {
	return func(c *BucketConfig) {
		c.Logger = logger
	}
}

// WithCapacity modifies the default capacity for a bucket.
// Returns a BucketConfigurator with the given capacity.
func WithCapacity(capacity int) BucketConfigurator {
	return func(c *BucketConfig) {
		c.Capacity = capacity
	}
}
