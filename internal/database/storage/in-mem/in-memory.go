package in_mem

import (
	"database/internal/database/storage"
	"fmt"
)

type InMemoryStorage struct {
	src map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		src: make(map[string]string),
	}
}

func (s *InMemoryStorage) Get(key string) storage.Result {
	val, ok := s.src[key]
	if !ok {
		return storage.Result{
			Out: fmt.Sprintf("[storage] value by key '%s' not found", key),
		}
	}

	return storage.Result{
		Out: val,
	}
}
func (s *InMemoryStorage) Set(key string, value string) storage.Result {
	s.src[key] = value

	return storage.Result{
		Out: "[storage] success execute set operation",
	}
}
func (s *InMemoryStorage) Delete(key string) storage.Result {
	delete(s.src, key)

	return storage.Result{
		Out: "[storage] success execute delete operation",
	}
}
