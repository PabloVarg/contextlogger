package contextlogger

import (
	"context"
	"log/slog"
)

func EmbedLoggingAttrs(ctx context.Context) context.Context {
	reservedCapacity := 20

	return context.WithValue(ctx, loggerKey, &handlerLogger{
		attrs: make(map[string]any, reservedCapacity),
	})
}

func WithGroup(ctx context.Context, name string) context.Context {
	reservedCapacity := 20

	logger, ok := loggerGroup(ctx, name)
	if !ok {
		logger = &handlerLogger{
			attrs: make(map[string]any, reservedCapacity),
		}
	}

	UpdateContext(ctx, name, logger)

	return context.WithValue(ctx, loggerKey, logger)
}

func loggerGroup(ctx context.Context, name string) (*handlerLogger, bool) {
	parent, ok := ctx.Value(loggerKey).(*handlerLogger)
	if !ok {
		return nil, false
	}

	value, ok := parent.attrs[name]
	if !ok {
		return nil, false
	}

	logger, ok := value.(*handlerLogger)
	if !ok {
		return nil, false
	}

	return logger, true
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
