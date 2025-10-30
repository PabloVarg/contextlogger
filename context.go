package contextlogger

import (
	"fmt"
	"log/slog"
	"maps"

	"github.com/google/uuid"
)

// Type to store values in context.Context.
type loggingContext = string

// Key to store logger struct in a context.Context.
var loggerKey = loggingContext("logger")

// Key to store slog.Logger to use for printing messages
var slogLoggerKey = loggingContext("slogLogger")

// handlerLogger holds key value pairs to log, it can hold other handlerLogger's
// to implement groups.
type handlerLogger struct {
	attrs map[string]any
}

// append stores into it's context: Pairs of [string, any] or slog.Attr's.
// For any other format, it will print the value or values given with a
// BADKEY name.
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

// asAttrs returns internal representation of context as slog.Attr's to be
// printed by a slog.Logger.
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
