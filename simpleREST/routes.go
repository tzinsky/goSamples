package main

import "net/http"

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
		Index,
	},
	Route{
		"TransactionCreate",
		"POST",
		"/transactions",
		TransactionCreate,
	},
	Route{
		"TransactionShow",
		"GET",
		"/transactions/{transactionId}",
		TransactionShow,
	},
	Route{
		"TransactionEdit",
		"PUT",
		"/transactions/{transactionId}",
		TransactionEdit,
	},
}
