package storage

import "errors"

// Define a simple not found error message
var (
	ErrNotFound = errors.New("Not Found")
)

// DB setsup the interface for a key/value pair memory storage construct
type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}
