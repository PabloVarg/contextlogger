package contextlogger

import "context"

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
