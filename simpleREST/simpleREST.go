package main

import (
    "net/http"
    "log"
)

func response (rw http.ResponseWriter, request *http.Request) {
    log.Printf("ACTION: %s - Remote: %s Local URI: %s",request.Method, 
                                             request.RemoteAddr,
                                             request.RequestURI)
    rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
    rw.Write([]byte ("Hello world!"))
}

func main() {
    http.HandleFunc("/", response)
    log.Fatal(http.ListenAndServe (":8080", nil))
}
