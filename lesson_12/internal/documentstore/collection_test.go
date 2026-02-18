package documentstore

import "testing"

func TestCollection_PutGetDeleteList(t *testing.T) {
	s := NewStore()
	col, err := s.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
	if err != nil {
		t.Fatalf("CreateCollection error: %v", err)
	}

	err = col.Put(Document{Fields: map[string]DocumentField{}})
	if err != ErrInvalidPrimaryKey {
		t.Fatalf("expected ErrInvalidPrimaryKey, got %v", err)
	}

	err = col.Put(Document{
		Fields: map[string]DocumentField{
			"id": {Type: DocumentFieldTypeNumber, Value: 1},
		},
	})
	if err != ErrInvalidPrimaryKey {
		t.Fatalf("expected ErrInvalidPrimaryKey, got %v", err)
	}

	err = col.Put(Document{
		Fields: map[string]DocumentField{
			"id": {Type: DocumentFieldTypeString, Value: ""},
		},
	})
	if err != ErrInvalidPrimaryKey {
		t.Fatalf("expected ErrInvalidPrimaryKey, got %v", err)
	}

	err = col.Put(Document{
		Fields: map[string]DocumentField{
			"id":   {Type: DocumentFieldTypeString, Value: "u1"},
			"name": {Type: DocumentFieldTypeString, Value: "Alina"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected Put error: %v", err)
	}

	got, err := col.Get("u1")
	if err != nil {
		t.Fatalf("unexpected Get error: %v", err)
	}
	if got == nil {
		t.Fatal("expected document, got nil")
	}
	if got.Fields["id"].Value != "u1" {
		t.Fatalf("expected id=u1, got %v", got.Fields["id"].Value)
	}

	if _, err := col.Get("missing"); err != ErrDocumentNotFound {
		t.Fatalf("expected ErrDocumentNotFound, got %v", err)
	}

	err = col.Put(Document{
		Fields: map[string]DocumentField{
			"id":   {Type: DocumentFieldTypeString, Value: "u1"},
			"name": {Type: DocumentFieldTypeString, Value: "AlinaUpdated"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected Put(update) error: %v", err)
	}

	list := col.List()
	if len(list) != 1 {
		t.Fatalf("expected list size 1, got %d", len(list))
	}

	if err := col.Delete("u1"); err != nil {
		t.Fatalf("unexpected Delete error: %v", err)
	}

	if err := col.Delete("u1"); err != ErrDocumentNotFound {
		t.Fatalf("expected ErrDocumentNotFound, got %v", err)
	}

	list = col.List()
	if len(list) != 0 {
		t.Fatalf("expected list size 0, got %d", len(list))
	}
}
