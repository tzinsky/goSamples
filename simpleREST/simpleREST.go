package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func response(rw http.ResponseWriter, request *http.Request) {
	log.Printf("ACTION: %s - Remote: %s Local URI: %s", request.Method,
		request.RemoteAddr,
		request.RequestURI)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Write([]byte("Hello world!"))
}

func main() {
	var handler http.HandlerFunc
	router := mux.NewRouter().StrictSlash(true)
	handler = Logger(response, "Index")
	router.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
