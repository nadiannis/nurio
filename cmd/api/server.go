package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *application) serve() error {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting API %s server on :%d\n", app.config.env, app.config.port)
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
