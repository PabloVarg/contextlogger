package contextlogger

import (
	"fmt"
	"log/slog"
	"maps"

	"github.com/google/uuid"
)

type loggingContext = string

var loggerKey = loggingContext("logger")

type handlerLogger struct {
	attrs map[string]any
}

func (l *handlerLogger) append(attrs ...any) {
	i := 0

	for {
		if i >= len(attrs) {
			return
		}

		// Handle string value pairs
		if key, ok := attrs[i].(string); ok && i+1 < len(attrs) {
			l.attrs[key] = attrs[i+1]
			i += 2
			continue
		}

		// Handle slog Attr's
		if attr, ok := attrs[i].(slog.Attr); ok {
			l.attrs[attr.Key] = attr.Value
			i++
			continue
		}

		// Handle bad formats
		l.attrs[fmt.Sprintf("!BADKEY:%s", uuid.NewString())] = attrs[i]
		i++
	}
}

func (l *handlerLogger) asAttrs() []any {
	attrs := make([]any, 0, len(l.attrs))
	for name, attr := range maps.All(l.attrs) {
		// Handle nested handlerLogger's
		if logger, ok := attr.(*handlerLogger); ok {
			attrs = append(attrs, slog.Group(name, logger.asAttrs()...))
			continue
		}

		// Handle values that are already slog.Attr's
		if attr, ok := attr.(slog.Attr); ok {
			attrs = append(attrs, attr)
			continue
		}

		attrs = append(attrs, slog.Any(name, attr))
	}

	return attrs
}
