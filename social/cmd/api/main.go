package main

import (
	"log"

	"github.com/agamyo168/social-blog/internal/env"
	"github.com/agamyo168/social-blog/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store: store,
	}
	
	log.Fatal(app.run(app.mount()))	
}