package contextlogger

import (
	"context"

	"github.com/pablovarg/contextlogger/buckets"
)

// loggerGroup retrieves a nested group in the current layer for context with a given name,
// if the key with name contains any other value or such key is not found, the bool response will be false,
func loggerGroup(ctx context.Context, name string) (*buckets.Bucket, bool) {
	parent, ok := ctx.Value(loggerKey).(*buckets.Bucket)
	if !ok {
		return nil, false
	}

	value, ok := parent.Attrs[name]
	if !ok {
		return nil, false
	}

	logger, ok := value.(*buckets.Bucket)
	if !ok {
		return nil, false
	}

	return logger, true
}
