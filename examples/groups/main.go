package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/pablovarg/contextlogger"
	"github.com/pablovarg/contextlogger/examples"
)

func main() {
	ctx := context.Background()

	defer func() {
		// Log added context, you usually want this to be called at the end of your process/action/handler,
		// thus usually this will be inside a defer statementt, but this can be called anywhere in your code.
		contextlogger.LogWithContext(ctx, slog.LevelInfo, "user information")
	}()

	ctx = contextlogger.EmbedLogger(ctx)
	processUser(contextlogger.WithGroup(ctx, "userProcess"))
	processTime(contextlogger.WithGroup(ctx, "timeProcess"))
}

func processUser(ctx context.Context) {
	user := examples.DefaultUser()

	// You can add key value pairs
	contextlogger.UpdateContext(ctx, "id", user.ID, "email", user.Email)

	// You can add slog.Attr's
	contextlogger.UpdateContext(ctx, slog.String("firstName", user.FirstName))
}

func processTime(ctx context.Context) {
	contextlogger.UpdateContext(ctx, "currentTime", time.Now())
}
