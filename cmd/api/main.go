package main

import (
	"backend-fundamentals/internal/env"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("BIND_PORT", ":8000"),
	}

	app := &application{
		config: cfg,
	}

	log.Printf("Server has started")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
