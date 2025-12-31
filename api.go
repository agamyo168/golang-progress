package main

import (
	"encoding/json"
	"net/http"
)
type api struct {
	addr string
}

var users = []User{}

//Go Receivers.
func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(
			 w, //Writer stream
			 err.Error(), //string
			 http.StatusInternalServerError) //error code
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request){
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
	}
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
}