package main

import "log"

func main() {
	cfg := config{
		addr: ":8000",
	}

	app := &application{
		config: cfg,
	}

	log.Printf("Server has started")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
