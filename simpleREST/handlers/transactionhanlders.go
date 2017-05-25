package handlers

import (
	"fmt"
	"net/http"
	"net/url"
)

func TransactionCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST - To be implemented soon - chill peoples!\n")

	/*
		Basic algorithm
			Assume accepted transaction
			Validate request
			if valid
				Create worker to make changes (deposit, transer, withdraw)
			else
				Indicate invalid transaction
			Return task for tracking and response

	*/

}

func TransactionIndex(w http.ResponseWriter, r *http.Request) {

	/*
		Basic algorithm
			Assume success (HTTP200)
			Get Parameters (count, filters)
			if Valid Parameters
				Obtain transactions in JSON array
			else
				Set error bad params
			return
	*/
	responseStatus := http.StatusOK
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	u, err := url.Parse(r.URL)
	params, _ := url.ParseQuery(u.RawQuery)

	w.WriteHeader(responseStatus)

}

func TransactionShow(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "GET - To be implemented soon - chill peoples!\n")

}

func TransactionEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PUT - To be implemented soon - chill peoples!\n")

}
