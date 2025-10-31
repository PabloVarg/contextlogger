package main

import (
	"context"
	"log/slog"

	"github.com/pablovarg/contextlogger"
	"github.com/pablovarg/contextlogger/examples"
)

func main() {
	ctx := context.Background()

	process(contextlogger.EmbedLoggingAttrs(ctx))
}

func process(ctx context.Context) {
	user := examples.DefaultUser()

	// You can add key value pairs
	contextlogger.UpdateContext(ctx, "id", user.ID, "email", user.Email)

	// You can add slog.Attr's
	contextlogger.UpdateContext(ctx, slog.String("firstName", user.FirstName))

	// Log added context
	contextlogger.LogWithContext(ctx, slog.LevelInfo, "user information")
}
