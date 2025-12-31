## Go http server routers

In Node.js terms, http.NewServeMux() is the equivalent of calling const app = express(). It creates a request router (or "multiplexer"). Its job is to look at the incoming HTTP request path and "multiplex" (match) it to the correct handler function.
|Concept |Express (Node.js) |Go (net/http)|
|-----------|-------------------|-------------|
|The Router| const app = express()| mux := http.NewServeMux()|
|Defining a Route| app.get('/home', callback)| mux.HandleFunc("/home", handler)|
|Starting the Server| app.listen(3000)| http.ListenAndServe(":3000", mux)|

### Request Handler

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
    // 1. Create the router (The "Express App")
	mux := http.NewServeMux()

    // 2. Define a route
	mux.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Go!")
	})

    // 3. Start the server using our mux
	http.ListenAndServe(":8080", mux)
}
```

### Global Variable DefaultServeMux

ou might see some Go tutorials that skip the mux := ... part and just write: http.HandleFunc("/path", handler)

This uses the DefaultServeMux (a global variable inside the http package). Why you should avoid the DefaultServeMux (and use your own mux):

    Security: Because it's a global variable, any third-party package you import could potentially register routes to your server without you knowing.

    Testing: It's much harder to test your code if everything is attached to a global state.

    Explicit is better: In Go philosophy, being explicit (creating your own mux) is almost always preferred over "magic" globals.
