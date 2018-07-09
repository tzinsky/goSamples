package main

import (
	"goSamples/servemux/handlers"
	"goSamples/servemux/storage"
	"log"
	"net/http"
)

func main() {

	db := storage.NewInMemoryDB()
	mux := http.NewServeMux()

	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.PutKey(db))
	log.Fatal(http.ListenAndServe(":8080", mux))

}
