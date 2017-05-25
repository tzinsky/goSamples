package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func parseUrlHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := url.Parse(r.RequestURI)
	params, _ := url.ParseQuery(u.RawQuery)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(params); err != nil {
		panic(err)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(parseUrlHandler))
	log.Fatal(http.ListenAndServe(":8080", mux))

}
