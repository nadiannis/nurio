package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /api/v1/healthcheck", http.HandlerFunc(app.healthcheck))

	return app.recoverPanic(app.requestLogger(mux))
}
