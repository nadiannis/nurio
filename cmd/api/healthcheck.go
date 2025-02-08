package main

import (
	"net/http"

	"github.com/nadiannis/nurio/internal/types"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	env := types.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     types.APIVersion,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
