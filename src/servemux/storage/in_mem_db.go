package storage

import "sync"

type inMemDb struct {
	m   map[string][]byte
	lck sync.RWMutex
}

// NewInMemoryDB - Creates a simple in-memory DB with locks to support concurrency
func NewInMemoryDB() DB {
	return &inMemDb{m: make(map[string][]byte)}
}

func (d *inMemDb) Get(key string) ([]byte, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	v, ok := d.m[key]

	if !ok {
		return nil, ErrNotFound
	}

	return v, nil

}

func (d *inMemDb) Set(key string, val []byte) error {
	d.lck.Lock()
	defer d.lck.Unlock()
	d.m[key] = val
	return nil

}
