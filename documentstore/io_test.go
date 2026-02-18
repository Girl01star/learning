package documentstore

import (
	"testing"
)

func TestStore_DumpAndRestore(t *testing.T) {
	s := NewStore()

	users, err := s.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
	if err != nil {
		t.Fatalf("CreateCollection error: %v", err)
	}

	err = users.Put(Document{
		Fields: map[string]DocumentField{
			"id":   {Type: DocumentFieldTypeString, Value: "u1"},
			"name": {Type: DocumentFieldTypeString, Value: "Alina"},
		},
	})
	if err != nil {
		t.Fatalf("Put error: %v", err)
	}

	dump, err := s.Dump()
	if err != nil {
		t.Fatalf("Dump error: %v", err)
	}
	if len(dump) == 0 {
		t.Fatalf("expected non-empty dump")
	}

	s2, err := NewStoreFromDump(dump)
	if err != nil {
		t.Fatalf("NewStoreFromDump error: %v", err)
	}

	users2, err := s2.GetCollection("users")
	if err != nil {
		t.Fatalf("GetCollection after restore error: %v", err)
	}

	got, err := users2.Get("u1")
	if err != nil {
		t.Fatalf("Get restored doc error: %v", err)
	}
	if got == nil {
		t.Fatalf("expected doc, got nil")
	}

	nameField, ok := got.Fields["name"]
	if !ok {
		t.Fatalf("expected field name")
	}
	if nameField.Type != DocumentFieldTypeString {
		t.Fatalf("expected name type %q, got %q", DocumentFieldTypeString, nameField.Type)
	}
	if nameField.Value != "Alina" {
		t.Fatalf("expected name value %q, got %v", "Alina", nameField.Value)
	}
}
