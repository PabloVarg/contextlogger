package contextlogger

import "context"

// loggerGroup retrieves a nested group in the current layer for context with a given name,
// if the key with name contains any other value or such key is not found, the bool response will be false,
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
