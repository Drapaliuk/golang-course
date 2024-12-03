package documentstore

import "testing"

func TestNewDocument(t *testing.T) {
	fields := map[string]DocumentField{
		"ID":   {Type: DocumentFieldTypeString, Value: "1"},
		"Name": {Type: DocumentFieldTypeString, Value: "Jack"},
	}
	doc := NewDocument(fields)

	if len(doc.Fields) != 2 {
		t.Fatalf("Expected 2 fields, got %d", len(doc.Fields))
	}
}
