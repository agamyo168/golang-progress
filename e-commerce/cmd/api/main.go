package main

import (
	"log/slog"
	"os"

	"github.com/agamyo168/e-commerce/internal/env"
)

func main(){
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn: env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/ecomm_db?sslmode=disable"),
		},
	}
	//Structured Log
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	app := &application{
		config: cfg,
	}
	if err := app.run(app.mount()); err != nil{
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}