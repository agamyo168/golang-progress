package main //Every go file must have a package declaration -- Go projects must have a main package - main file and main function where the function is the entrypoint.

import (
	"log"
	"net/http"
)


type server struct {
	addr string
}
//Go Receiver.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
			case "/":
				w.Write([]byte("index page"))
				return
			case "/users":
				w.Write([]byte("users page"))
				return
		}
		
	default: 
		w.Write([]byte("404 page"))
	}
}

func main ()  {
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}