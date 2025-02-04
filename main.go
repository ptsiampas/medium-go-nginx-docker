package main

import (
	"flag"
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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",*port), router))
}

func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, i'm a golang microservice")
}
