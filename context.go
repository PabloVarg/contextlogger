package contextlogger

// Type to store values in context.Context.
type loggingContext = string

// Key to store logger struct in a context.Context.
var loggerKey = loggingContext("logger")

// Key to store slog.Logger to use for printing messages
var slogLoggerKey = loggingContext("slogLogger")
