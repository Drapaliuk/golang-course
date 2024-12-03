package documentstore

import (
	"testing"
)

func TestNewStore(t *testing.T) {
	store := NewStore()

	if len(store.Collections) != 0 {
		t.Fatalf("Expected empty collections, got %d", len(store.Collections))
	}
}

func TestCreateCollection(t *testing.T) {
	store := NewStore()
	cfg := CollectionConfig{Name: "testCollection", PrimaryKey: "ID"}

	collection, err := store.CreateCollection(cfg)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if collection.Configs.Name != "testCollection" {
		t.Fatalf("Expected collection name 'testCollection', got %v", collection.Configs.Name)
	}
}

func TestGetCollection(t *testing.T) {
	store := NewStore()
	cfg := CollectionConfig{Name: "testCollection", PrimaryKey: "ID"}
	store.CreateCollection(cfg)

	collection, err := store.GetCollection("testCollection")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if collection == nil {
		t.Fatal("Expected to find the collection, but got nil")
	}
}

func TestDeleteCollection(t *testing.T) {
	store := NewStore()
	cfg := CollectionConfig{Name: "testCollection", PrimaryKey: "ID"}
	store.CreateCollection(cfg)

	err := store.DeleteCollection("testCollection")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(store.Collections) != 0 {
		t.Fatalf("Expected no collections left, got %d", len(store.Collections))
	}
}
