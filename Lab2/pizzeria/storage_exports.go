package pizzeria

import "lab2/pizzeria/storage"

type (
	Storage       = storage.Storage
	MemoryStorage = storage.MemoryStorage
)

var (
	NewMemoryStorage = storage.NewMemoryStorage
)
