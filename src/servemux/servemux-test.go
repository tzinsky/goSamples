package main

import (
	"log"
	"net/http"
	"servemux/handlers"
	"servemux/storage"
)

func main() {

	db := storage.NewInMemoryDB()
	mux := http.NewServeMux()

	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.PutKey(db))
	log.Fatal(http.ListenAndServe(":8080", mux))

}
