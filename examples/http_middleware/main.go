package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/pablovarg/contextlogger"
	"github.com/pablovarg/contextlogger/examples"
	"github.com/pablovarg/contextlogger/middlewares/httpmiddleware"
)

func main() {
	runServer()
}

func runServer() {
	handlerWithMiddleware := httpmiddleware.LoggingMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := examples.DefaultUser()

			contextlogger.UpdateContext(
				contextlogger.WithGroup(r.Context(), "userInfo"),
				"id", user.ID,
				"email", user.Email,
				"firstName", user.FirstName,
			)
		}),
		httpmiddleware.WithLogger(slog.New(slog.NewJSONHandler(os.Stdout, nil))),
	)

	http.DefaultServeMux.Handle("GET /users/info", handlerWithMiddleware)
	http.ListenAndServe(":8000", http.DefaultServeMux)
}
