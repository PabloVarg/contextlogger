# Go's Context Logger

<img align="right" src="assets/contextlogger.svg">[![Releases](https://img.shields.io/github/v/release/pablovarg/contextlogger)](https://github.com/pablovarg/contextlogger/releases)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/pablovarg/contextlogger)
![GitHub License](https://img.shields.io/github/license/pablovarg/contextlogger)
![GitHub Tag](https://img.shields.io/github/v/tag/pablovarg/contextlogger)

ContextLogger implements contextual logging using `slog` by embedding itself into `context.Context`'s.
Making it effortless to maintain consistent log context throughout your application's call stack.

It has the following features:

- Context-Embedded Logging - Logger travels with your context.Context, eliminating manual logger passing
- slog Familiarity - Built on top of Go's standard log/slog package with a familiar API
- Structured Logging - Support for key-value pairs and nested groups for organized log output
- Highly Configurable - Customize log levels, handlers, formatters, and output destinations
- Middleware Ready - Built-in HTTP middleware for automatic request logging, and more coming soon!
- Multiple Handlers - Support for JSON, text, and custom log formats. Basically, everything that slog supports
- Performance Focused - Minimal overhead with efficient context propagation

## Why use Context Logger?

Traditional logging often requires passing a logger instance through your entire call stack or using a global logger that lacks request-specific context. Context Logger solves this by:

1. Embedding the logger in context - Access your logger anywhere you have a context.Context
2. Automatic context accumulation - Build up contextual information as requests flow through layers
3. Cleaner function signatures - No need to pass logger as a parameter to every function

## Installation

Use go get:

```sh
go get github.com/pablovarg/contextlogger
```

Then import the contextlogger package into your code:

```sh
import "github.com/pablovarg/contextlogger"
```

## Usage

```go
ctx := context.Background()

// Embed a logger into the context
ctx = contextlogger.EmbedLogger(ctx)

// Update the stored context
contextlogger.UpdateContext(ctx, "id", user.ID, "email", user.Email)

// Create groups for different parts of your code,
// and update the context as above inside this group
processUser(contextlogger.WithGroup(ctx, "userProcess"))

// To print the whole accumulated context
contextlogger.LogWithContext(ctx, slog.LevelInfo, "user information")
```

### Examples

- [Basic](https://github.com/PabloVarg/contextlogger/blob/main/examples/basic/main.go)
- [Custom Logger](https://github.com/PabloVarg/contextlogger/blob/main/examples/custom_logger/main.go)
- [Groups](https://github.com/PabloVarg/contextlogger/blob/main/examples/groups/main.go)
- [HTTP Middleware](https://github.com/PabloVarg/contextlogger/blob/main/examples/http_middleware/main.go)

## Contributing

This project is currently open to contributions from the community, some things you can work on:

- Support for different logging libraries
- Include a middleware for your favorite framework
- Anything else you can imagine
