package main

import (
	"log"
	"net/http"
	"time"

	"github.com/agamyo168/e-commerce/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//This is where we inject our dependecies!!!
type application struct {
	config config
	//logger
}
//mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//Good for rate-limiting -> remember the C# project!
	r.Use(middleware.RealIP) //Analytics and tracing
	r.Use(middleware.RequestID) //nice for tracing logs 
	//Time out request if not finished in 60 seconds
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type","application/json")
		w.Write([]byte("ok"))
	})
	productService := products.NewService()
	productHandler := products.NewHandler(productService)
	r.Get("/v1/products", productHandler.ListProducts)

	return  r;
} 
//run
func (app *application) run(h http.Handler) error {
		srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30, 
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	log.Printf("server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}
