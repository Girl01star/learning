package documentstore

import (
	"reflect"
	"testing"
)

func TestDocumentFieldType_Constants(t *testing.T) {
	if DocumentFieldTypeString != "string" {
		t.Fatalf("expected DocumentFieldTypeString = %q, got %q", "string", DocumentFieldTypeString)
	}
	if DocumentFieldTypeNumber != "number" {
		t.Fatalf("expected DocumentFieldTypeNumber = %q, got %q", "number", DocumentFieldTypeNumber)
	}
	if DocumentFieldTypeBool != "bool" {
		t.Fatalf("expected DocumentFieldTypeBool = %q, got %q", "bool", DocumentFieldTypeBool)
	}
	if DocumentFieldTypeArray != "array" {
		t.Fatalf("expected DocumentFieldTypeArray = %q, got %q", "array", DocumentFieldTypeArray)
	}
	if DocumentFieldTypeObject != "object" {
		t.Fatalf("expected DocumentFieldTypeObject = %q, got %q", "object", DocumentFieldTypeObject)
	}
}

func TestDocument_ZeroValue(t *testing.T) {
	var d Document

	if d.Fields != nil {
		t.Fatalf("expected zero-value Document.Fields to be nil, got %#v", d.Fields)
	}
}

func TestDocument_CreateAndReadFields(t *testing.T) {
	d := Document{
		Fields: map[string]DocumentField{
			"id":   {Type: DocumentFieldTypeString, Value: "u1"},
			"age":  {Type: DocumentFieldTypeNumber, Value: float64(25)},
			"vip":  {Type: DocumentFieldTypeBool, Value: true},
			"tags": {Type: DocumentFieldTypeArray, Value: []any{"a", "b"}},
			"meta": {Type: DocumentFieldTypeObject, Value: map[string]any{"city": "Kharkiv"}},
		},
	}

	f, ok := d.Fields["id"]
	if !ok {
		t.Fatalf("expected field %q to exist", "id")
	}
	if f.Type != DocumentFieldTypeString {
		t.Fatalf("expected field %q type %q, got %q", "id", DocumentFieldTypeString, f.Type)
	}
	if f.Value != "u1" {
		t.Fatalf("expected field %q value %v, got %v", "id", "u1", f.Value)
	}

	ft, ok := d.Fields["tags"]
	if !ok {
		t.Fatalf("expected field %q to exist", "tags")
	}
	if ft.Type != DocumentFieldTypeArray {
		t.Fatalf("expected field %q type %q, got %q", "tags", DocumentFieldTypeArray, ft.Type)
	}
	arr, ok := ft.Value.([]any)
	if !ok {
		t.Fatalf("expected field %q value to be []any, got %T", "tags", ft.Value)
	}
	if !reflect.DeepEqual(arr, []any{"a", "b"}) {
		t.Fatalf("expected tags %v, got %v", []any{"a", "b"}, arr)
	}

	fm, ok := d.Fields["meta"]
	if !ok {
		t.Fatalf("expected field %q to exist", "meta")
	}
	if fm.Type != DocumentFieldTypeObject {
		t.Fatalf("expected field %q type %q, got %q", "meta", DocumentFieldTypeObject, fm.Type)
	}
	obj, ok := fm.Value.(map[string]any)
	if !ok {
		t.Fatalf("expected field %q value to be map[string]any, got %T", "meta", fm.Value)
	}
	if obj["city"] != "Kharkiv" {
		t.Fatalf("expected meta.city = %q, got %v", "Kharkiv", obj["city"])
	}
}
