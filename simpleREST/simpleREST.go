package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Print("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))
	})
}

func response(rw http.ResponseWriter, request *http.Request) {
	log.Printf("ACTION: %s - Remote: %s Local URI: %s", request.Method,
		request.RemoteAddr,
		request.RequestURI)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Write([]byte("Hello world!"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", response)
	log.Fatal(http.ListenAndServe(":8080", router))
}
