package documentstore

import "testing"

func TestNewStore(t *testing.T) {
	s := NewStore()
	if s == nil {
		t.Fatalf("expected store, got nil")
	}
	if s.collections == nil {
		t.Fatalf("expected collections map to be initialized")
	}
}

func TestStore_CreateGetDeleteCollection(t *testing.T) {
	s := NewStore()

	_, err := s.CreateCollection("", &CollectionConfig{PrimaryKey: "id"})
	if err != ErrInvalidCollectionName {
		t.Fatalf("expected ErrInvalidCollectionName, got %v", err)
	}

	col, err := s.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if col == nil {
		t.Fatalf("expected collection, got nil")
	}

	_, err = s.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
	if err != ErrCollectionAlreadyExists {
		t.Fatalf("expected ErrCollectionAlreadyExists, got %v", err)
	}

	got, err := s.GetCollection("users")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got == nil {
		t.Fatalf("expected collection, got nil")
	}

	_, err = s.GetCollection("nope")
	if err != ErrCollectionNotFound {
		t.Fatalf("expected ErrCollectionNotFound, got %v", err)
	}

	err = s.DeleteCollection("users")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	err = s.DeleteCollection("users")
	if err != ErrCollectionNotFound {
		t.Fatalf("expected ErrCollectionNotFound, got %v", err)
	}
}
