```bash
go get -u github.com/go-chi/chi/v5
```

Instead of using mux the standard way to make request handlers we will use chi (router library) for that.
It only solves one problem: “How do I get this URL to this function?”

## Why we use it (The "Gap" it fills)

The Go standard library (net/http) actually got a huge upgrade recently (Go 1.22), so it can handle basic routing quite well. However, we still use Chi for three main "quality of life" reasons:

    Middleware Grouping: This is the killer feature. You can say "this group of routes needs Auth, but this group doesn't."
    Go

```go
r.Route("/admin", func(r chi.Router) {
    r.Use(AdminOnlyMiddleware)
    r.Get("/dashboard", app.dashboardHandler)
})
```

Clean URL Params: Chi makes it incredibly easy to grab variables from the URL.
Chi: `/users/{userID} → chi.URLParam(r, "userID")`

Standard Signature: Because it uses http.HandlerFunc, any middleware written by the Go community—ever—will work with Chi. You aren't locked into a "Chi-only" ecosystem.

###

middleware.Recoverer is a piece of middleware provided by Chi that gracefully handles panics. Instead of letting the entire server crash, it:

    Intercepts (Recovers) the panic.

    Logs the stack trace so you can see exactly where the code failed.

    Sends a 500 Internal Server Error back to the client.

    Keeps the server alive to handle other incoming requests.
