package documentstore

import (
	"errors"
	"fmt"
	u "lesson_3/utils"
)

type DocumentFieldType string

var store = []Document{}

var ErrDocKeyNotProvided = errors.New("document key not provided")
var ErrDocKeyIsNotString = errors.New("document key must be a string")

const FieldKeyName = "key"
const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value any
}

type Document struct {
	Fields map[string]DocumentField
}

func NewDocument(key string) Document {
	return Document{
		Fields: map[string]DocumentField{
			FieldKeyName: {Type: DocumentFieldTypeString, Value: key},
		},
	}
}

func Put(doc Document) error {
	keyField, ok := doc.Fields[FieldKeyName]

	if !ok {
		return ErrDocKeyNotProvided
	} else if keyField.Type != DocumentFieldTypeString {
		return ErrDocKeyIsNotString
	}

	isKeyExists := u.Some(store, func(storedDoc Document, _ int) bool {
		return storedDoc.Fields[FieldKeyName].Value == doc.Fields[FieldKeyName].Value
	})

	if isKeyExists {
		return fmt.Errorf("document with key %v already exists", doc.Fields[FieldKeyName].Value)
	}

	store = append(store, doc)

	return nil
}

func Get(key string) (bool, *Document) {
	docPointer := u.Find(store, func(doc Document, _ int) bool {
		return doc.Fields[FieldKeyName].Value == key
	})

	if docPointer != nil {
		return true, docPointer
	} else {
		return false, nil
	}
}

func Delete(key string) bool {
	targetDocIdx := u.FindIndex(store, func(doc Document, _ int) bool {
		return doc.Fields[FieldKeyName].Value == key
	})

	if targetDocIdx == -1 {
		return false
	}

	store = u.Delete(store, targetDocIdx)
	return true
}

func List() []Document {
	return store
}
