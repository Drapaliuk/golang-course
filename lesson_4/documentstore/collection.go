package documentstore

import (
	"fmt"
	u "lesson_3/utils"
)

// Collection - дає можливість зберігати ат отримувати документи.
// У кожної колекції можна задати (через конфіг) яке саме поле у документа вважати primary ключем.

type Collection struct {
	Configs   CollectionConfig
	Documents []Document
}

func NewCollection(cfg CollectionConfig) *Collection {
	return &Collection{
		Configs: cfg,
	}
}

type CollectionConfig struct {
	PrimaryKey string
	Name       string
}

func (s *Collection) Put(doc Document) error {
	// if doc[s.Configs.PrimaryKey] {
	// 	// як?
	// }

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

func (s *Collection) Get(key string) (bool, *Document) {
	docPointer := u.Find(store, func(doc Document, _ int) bool {
		return doc.key == key
	})

	if docPointer != nil {
		return true, docPointer
	} else {
		return false, nil
	}
}

func (s *Collection) Delete(key string) bool {
	targetDocIdx := u.FindIndex(store, func(doc Document, _ int) bool {
		return doc.key == key
	})

	if targetDocIdx == -1 {
		return false
	}

	u.Delete(store, targetDocIdx)
	return true
}

func (s *Collection) List() []Document {
	return s.Documents
}
