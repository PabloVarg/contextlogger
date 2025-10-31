package contextlogger

import (
	"context"
	"log/slog"

	"github.com/pablovarg/contextlogger/internal/buckets"
)

// EmbedLoggingAttrs embeds a contextual logger into the passed context.Context.
// Returns a new context.Context with the embeded logger.
func EmbedLoggingAttrs(
	ctx context.Context,
	logger *slog.Logger,
	opts ...buckets.BucketConfigurator,
) context.Context {
	conf := buckets.DefaultBucketConfig()
	for _, opt := range opts {
		opt(conf)
	}

	ctx = context.WithValue(ctx, loggerKey, &buckets.Bucket{
		Attrs: make(map[string]any, conf.Capacity),
	})
	ctx = context.WithValue(ctx, slogLoggerKey, logger)

	return ctx
}

// WithGroup creates a group with the given name on the current group give by the
// context.Context passed into the function. Returns inner context.Context created
// for the new group.
func WithGroup(
	ctx context.Context,
	name string,
	opts ...buckets.BucketConfigurator,
) context.Context {
	conf := buckets.DefaultBucketConfig()
	for _, opt := range opts {
		opt(conf)
	}

	logger, ok := loggerGroup(ctx, name)
	if !ok {
		logger = &buckets.Bucket{
			Attrs: make(map[string]any, conf.Capacity),
		}
	}

	UpdateContext(ctx, name, logger)

	return context.WithValue(ctx, loggerKey, logger)
}

// UpdateContext updates a series of [key, values] from the current logging context
// the input is the same as slog args, i.e.: pairs of [string, any] or slog.Attr's.
// Malformed inputs that don't conform to this, will be assigned to a BADKEY key.
func UpdateContext(ctx context.Context, attrs ...any) {
	logger, ok := ctx.Value(loggerKey).(*buckets.Bucket)
	if !ok {
		return
	}

	logger.Append(attrs...)
}

// LogWithContext logs everything stored into the context by calls to UpdateContext
func LogWithContext(ctx context.Context, level slog.Level, msg string) {
	handlerLogger, ok := ctx.Value(loggerKey).(*buckets.Bucket)
	if !ok {
		return
	}

	logger, ok := ctx.Value(slogLoggerKey).(*slog.Logger)
	if !ok {
		return
	}

	logger.Log(ctx, level, msg, handlerLogger.AsAttrs()...)
}
