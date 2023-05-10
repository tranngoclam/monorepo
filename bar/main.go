package main

import (
	"fmt"
	"lib/pkg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		message := fmt.Sprintf("Hello world from 'bar'! 1 + 2 = %d", pkg.SumInt(1, 2))
		_, _ = w.Write([]byte(message))
	})

	log.Println("starting server at http://localhost:8080/")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
