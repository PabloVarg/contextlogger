package contextlogger

import (
	"context"
	"log/slog"
)

func EmbedLoggingAttrs(ctx context.Context, opts ...BucketConfigurator) context.Context {
	conf := defaultBucketConfig()
	for _, opt := range opts {
		opt(conf)
	}

	return context.WithValue(ctx, loggerKey, &handlerLogger{
		attrs: make(map[string]any, conf.capacity),
	})
}

func WithGroup(ctx context.Context, name string, opts ...BucketConfigurator) context.Context {
	conf := defaultBucketConfig()
	for _, opt := range opts {
		opt(conf)
	}

	logger, ok := loggerGroup(ctx, name)
	if !ok {
		logger = &handlerLogger{
			attrs: make(map[string]any, conf.capacity),
		}
	}

	UpdateContext(ctx, name, logger)

	return context.WithValue(ctx, loggerKey, logger)
}

func UpdateContext(ctx context.Context, attrs ...any) {
	logger, ok := ctx.Value(loggerKey).(*handlerLogger)
	if !ok {
		return
	}

	logger.append(attrs...)
}

func LogWithContext(ctx context.Context, logger *slog.Logger, level slog.Level, msg string) {
	handlerLogger, ok := ctx.Value(loggerKey).(*handlerLogger)
	if !ok {
		return
	}

	logger.Log(ctx, level, msg, handlerLogger.asAttrs()...)
}
