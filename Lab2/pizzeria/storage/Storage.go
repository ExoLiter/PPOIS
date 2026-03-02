package storage

import (
	"fmt"
	"sync"
)

type Storage interface {
	Save(key string, value any) error
	Load(key string) (any, error)
}

type MemoryStorage struct {
	data map[string]any
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]any)}
}

func (m *MemoryStorage) Save(key string, value any) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Load(key string) (any, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.data[key]
	if !ok {
		return nil, fmt.Errorf("storage: %s not found", key)
	}
	return v, nil
}
