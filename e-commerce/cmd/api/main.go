package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/agamyo168/e-commerce/internal/env"
	"github.com/jackc/pgx/v5"
)

func main(){
	ctx := context.Background()
	//Structured Log
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "postgres://admin:adminpassword@localhost/ecomm_db?sslmode=disable"),
		},
	}
	//db connection
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic("Unable to connect to database")
	}
	defer conn.Close(ctx)

	app := &application{
		config: cfg,
	}
	if err := app.run(app.mount()); err != nil{
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}