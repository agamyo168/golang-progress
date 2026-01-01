package main

import (
	"log"

	"github.com/agamyo168/social-blog/internal/env"
	"github.com/agamyo168/social-blog/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://username@adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store: store,
	}
	
	log.Fatal(app.run(app.mount()))	
}