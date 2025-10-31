package buckets

import (
	"fmt"
	"log/slog"
	"maps"

	"github.com/google/uuid"
)

// Bucket holds key value pairs to log, it can hold other Bucket's
// to implement groups.
type Bucket struct {
	Attrs map[string]any
}

// append stores into it's context: Pairs of [string, any] or slog.Attr's.
// For any other format, it will print the value or values given with a
// BADKEY name.
func (l *Bucket) Append(attrs ...any) {
	i := 0

	for {
		if i >= len(attrs) {
			return
		}

		// Handle string value pairs
		if key, ok := attrs[i].(string); ok && i+1 < len(attrs) {
			l.Attrs[key] = attrs[i+1]
			i += 2
			continue
		}

		// Handle slog Attr's
		if attr, ok := attrs[i].(slog.Attr); ok {
			l.Attrs[attr.Key] = attr.Value
			i++
			continue
		}

		// Handle bad formats
		l.Attrs[fmt.Sprintf("!BADKEY:%s", uuid.NewString())] = attrs[i]
		i++
	}
}

// asAttrs returns internal representation of context as slog.Attr's to be
// printed by a slog.Logger.
func (l *Bucket) AsAttrs() []any {
	attrs := make([]any, 0, len(l.Attrs))
	for name, attr := range maps.All(l.Attrs) {
		// Handle nested handlerLogger's
		if logger, ok := attr.(*Bucket); ok {
			attrs = append(attrs, slog.Group(name, logger.AsAttrs()...))
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
