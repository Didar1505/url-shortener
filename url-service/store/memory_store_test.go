package store

import "testing"

func TestMemoryStore_SaveAndGet(t *testing.T) {
	store := NewMemoryStore()

	err := store.Save("abc123", "https://example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	gotURL, found := store.Get("abc123")
	if !found {
		t.Fatal("expected code to be found")
	}

	if gotURL != "https://example.com" {
		t.Fatalf("expected URL %q, got %q", "https://example.com", gotURL)
	}
}

func TestMemoryStore_GetNotFound(t *testing.T) {
	store := NewMemoryStore()

	_, found := store.Get("missing")
	if found {
		t.Fatal("expected code to be missing")
	}
}

func TestMemoryStore_SaveDuplicateCode(t *testing.T) {
	store := NewMemoryStore()

	err := store.Save("abc123", "https://example.com")
	if err != nil {
		t.Fatalf("expected no error on first save, got %v", err)
	}

	err = store.Save("abc123", "https://another-example.com")
	if err == nil {
		t.Fatal("expected error on duplicate code, got nil")
	}
}