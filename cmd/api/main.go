package main

import (
	"flag"
	"log"
)

const version = "0.1.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")

	flag.Parse()

	app := &application{
		config: cfg,
	}

	log.Fatal(app.serve())
}
