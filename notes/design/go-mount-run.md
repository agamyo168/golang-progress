The mount() function’s only job is to define the "API surface area." It maps URLs to logic.
The run() function handles the "Plumbing." It manages timeouts, SSL certificates, port numbers, and the actual start-up of the process.

```go
func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	return mux
}
func (app *application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30, //Max is 30 seconds for writing a response.
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	log.Printf("server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
```

## Testability

This is the biggest win. If you want to write a Unit Test for an endpoint, you don't want to actually start a real web server that opens a network port.

Because mount() returns a \*http.ServeMux, you can pass that mux into httptest.NewServer. This allows you to simulate requests in your tests without ever calling run().
