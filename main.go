package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := flag.Int("port", 8080, "Port to run on")
	flag.Parse()
	
	router := mux.NewRouter()
	router.HandleFunc("/", DoHealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, i'm a golang microservice")
}
