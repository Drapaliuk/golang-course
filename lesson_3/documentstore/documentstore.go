package documentstore

import (
	"fmt"
	u "lesson_3/utils"
)

type DocumentFieldType string

var store = []Document{}

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type DocumentFieldType
}

type Document struct {
	key    string
	Fields map[string]DocumentField
}

func NewDocument(key string) Document {
	return Document{
		key: key,
	}
}

func Put(doc Document) error {
	if len(doc.key) == 0 {
		return fmt.Errorf("Error: Document must have field \"key\"")
	}

	isKeyExists := u.Some(store, func(storedDoc Document, _ int) bool {
		return storedDoc.key == doc.key
	})

	if isKeyExists {
		return fmt.Errorf("Error: Document with key %d exists already", doc.key)
	}

	store = append(store, doc)

	return nil
}

func Get(key string) (bool, *Document) {
	docPointer := u.Find(store, func(doc Document, _ int) bool {
		return doc.key == key
	})

	if docPointer != nil {
		return true, docPointer
	} else {
		return false, nil
	}
}

func Delete(key string) bool {
	targetDocIdx := u.FindIndex(store, func(doc Document, _ int) bool {
		return doc.key == key
	})

	if targetDocIdx == -1 {
		return false
	}

	u.Del(store, targetDocIdx)
	return true
}

func List() []Document {
	// А якщо я віддам поінтер на приватну змінну через геттер, далі той хто отримає зможе робити з нею, все що завгодно?

	return store
}
