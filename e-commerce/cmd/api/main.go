package main

import (
	"log/slog"
	"os"
)

func main(){
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: "postgres://admin:adminpassword@local/ecommerce?sslmode=disable",
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
		os.Exit(1) //Why not Panic?
	}
}