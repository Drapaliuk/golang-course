package documentstore

import (
	"testing"
)

func TestPutDocument(t *testing.T) {
	store := NewStore()
	cfg := CollectionConfig{Name: "testCollection", PrimaryKey: "ID"}
	collection, _ := store.CreateCollection(cfg)

	doc := Document{
		Fields: map[string]DocumentField{
			"ID":   {Type: DocumentFieldTypeString, Value: "1"},
			"Name": {Type: DocumentFieldTypeString, Value: "John"},
		},
	}

	err := collection.Put(doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(collection.Documents) != 1 {
		t.Fatalf("Expected 1 document, got %d", len(collection.Documents))
	}
}

func TestGetDocument(t *testing.T) {
	store := NewStore()
	cfg := CollectionConfig{Name: "testCollection", PrimaryKey: "ID"}
	collection, _ := store.CreateCollection(cfg)

	doc := Document{
		Fields: map[string]DocumentField{
			"ID":   {Type: DocumentFieldTypeString, Value: "1"},
			"Name": {Type: DocumentFieldTypeString, Value: "John"},
		},
	}

	collection.Put(doc)

	retrievedDoc, err := collection.Get("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if retrievedDoc == nil {
		t.Fatal("Expected to retrieve the document, but got nil")
	}
	if retrievedDoc.Fields["Name"].Value != "John" {
		t.Fatalf("Expected Name 'John', got %v", retrievedDoc.Fields["Name"].Value)
	}
}

func TestDeleteDocument(t *testing.T) {
	store := NewStore()
	cfg := CollectionConfig{Name: "testCollection", PrimaryKey: "ID"}
	collection, _ := store.CreateCollection(cfg)

	doc := Document{
		Fields: map[string]DocumentField{
			"ID":   {Type: DocumentFieldTypeString, Value: "1"},
			"Name": {Type: DocumentFieldTypeString, Value: "John"},
		},
	}

	collection.Put(doc)
	err := collection.Delete("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(collection.Documents) != 0 {
		t.Fatalf("Expected 0 documents, got %d", len(collection.Documents))
	}
}
