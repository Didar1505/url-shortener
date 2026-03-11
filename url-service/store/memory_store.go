package store

import (
	"errors"
	"sync"
)

var ErrCodeAlreadyExists = errors.New("short code already exists")

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (ms *MemoryStore) Save(code string, originalURL string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.data[code]; ok {
		return ErrCodeAlreadyExists
	}
	
	ms.data[code] = originalURL
	return nil
}

func (ms *MemoryStore) Get(code string) (string, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	url, exist := ms.data[code]
	
	return url, exist
}

func (ms *MemoryStore) Count() int {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	return len(ms.data)
}
