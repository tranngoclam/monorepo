package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Hello world from 'bar'!`))
	})

	log.Println("starting server at http://localhost:8080/")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
