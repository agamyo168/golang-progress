package main //Every go file must have a package declaration -- Go projects must have a main package (you can name it anything) - main file and main function where the function is the entrypoint.

import (
	"net/http"
)

func main ()  {

	a := &api{addr: ":8080"}
	mux := http.NewServeMux()
	
	//Server is a struct we can configure and then call ListenAndServe instead of passing everything as an argument and gives us more options to configure.
	srv := &http.Server{
		Addr: a.addr,
		Handler: mux,
	}
	mux.HandleFunc("GET /users", a.getUsersHandler)
	mux.HandleFunc("POST /users", a.createUsersHandler)

	if err := srv.ListenAndServe();  err !=nil {
		panic(err)
	}

}