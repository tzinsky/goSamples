package main

import (
	"time"
)

type Transaction struct {
	Number     int       `json:"name"`
	Payee      string    `json:"payee"`
	Amount     float64   `json:"amount"`
	Reconciled bool      `json:"reconsiled"`
	Date       time.Time `json:"date"`
	Note       string    `json:"note"`
}

// The list of transactions
type Transactions []Transaction
