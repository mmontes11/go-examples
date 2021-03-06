package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port     = ":8080"
	endpoint = "/api"
)

func main() {
	newAPI()
}

func newAPI() {
	r := mux.NewRouter()
	r.HandleFunc(endpoint, handleRequest)
	log.Println("Server listening at", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界\n"))
}
