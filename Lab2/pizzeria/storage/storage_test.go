package storage_test

import (
	"testing"

	"lab2/pizzeria/storage"
)

func TestMemoryStorage(t *testing.T) {
	s := storage.NewMemoryStorage()
	if err := s.Save("order", 1); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	value, err := s.Load("order")
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if value.(int) != 1 {
		t.Fatalf("unexpected value %v", value)
	}
	if _, err := s.Load("missing"); err == nil {
		t.Fatalf("expected missing value error")
	}
}
