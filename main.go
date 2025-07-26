package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users handler"))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users handler"))
}

func main() {
	api := &api{addr: ":8081"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
