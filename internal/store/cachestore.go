package store

import (
	"bookstore/store"
	"sync"
)

type CacheStore struct {
	sync.RWMutex
	books map[string]*store.Book
}
