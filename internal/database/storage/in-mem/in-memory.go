package in_mem

import (
	"database/internal/database/storage"
	"fmt"
	"sync"
)

type InMemoryStorage struct {
	mu  sync.Mutex
	src map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		src: make(map[string]string),
	}
}

func (s *InMemoryStorage) get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.src[key]

	return val, ok
}

func (s *InMemoryStorage) Get(key string) storage.Result {
	val, ok := s.get(key)

	if !ok {
		return storage.Result{
			Out: fmt.Sprintf("value by key '%s' not found", key),
		}
	}

	return storage.Result{
		Out: val,
	}
}
func (s *InMemoryStorage) Set(key string, value string) storage.Result {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.src[key] = value

	return storage.Result{
		Out: "OK",
	}
}
func (s *InMemoryStorage) Delete(key string) storage.Result {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.src, key)

	return storage.Result{
		Out: "OK",
	}
}
