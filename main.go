package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
		case "/users":
			w.Write([]byte("users page"))
		default:
			w.Write([]byte("Not found"))
		}
	default:
		w.Write([]byte("404 name"))

	}
}

func main() {
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}
