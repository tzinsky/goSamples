package main

import "net/http"
import "goSamples/simpleREST/handlers"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"TransactionCreate",
		"POST",
		"/transactions",
		handlers.TransactionCreate,
	},
	Route{
		"TransactionShow",
		"GET",
		"/transactions/{transactionId}",
		handlers.TransactionShow,
	},
	Route{
		"TransactionEdit",
		"PUT",
		"/transactions/{transactionId}",
		handlers.TransactionEdit,
	},
}
