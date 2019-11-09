package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	connectToDb()
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./frontend")))

	log.Fatal(http.ListenAndServe(":8080", r))
}
