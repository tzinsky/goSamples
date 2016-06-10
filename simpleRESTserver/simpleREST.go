package main

import (
    "net/http"
    "log"
)

func response (rw http.ResponseWriter, request *http.Request) {
    rw.Write([]byte ("Hello world!"))
}

func main() {
    http.HandleFunc("/", response)
    log.Fatal(http.ListenAndServe (":8080", nil))
}
