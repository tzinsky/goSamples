package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func response (rw http.ResponseWriter, request *http.Request) {
    log.Printf("ACTION: %s - Remote: %s Local URI: %s",request.Method, 
                                             request.RemoteAddr,
                                             request.RequestURI)
    rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
    rw.Write([]byte ("Hello world!"))
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", response)
    log.Fatal(http.ListenAndServe (":8080", router))
}
