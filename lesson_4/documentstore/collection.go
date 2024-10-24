package documentstore

import (
	"fmt"
	u "lesson_4/utils"
)

type Collection struct {
	Configs   CollectionConfig
	Documents []*Document
}

func NewCollection(cfg CollectionConfig) *Collection {
	return &Collection{
		Configs: cfg,
	}
}

type CollectionConfig struct {
	PrimaryKey string // зробити як дженерик
	Name       string
}

func (s *Collection) GetPk() string {
	return s.Configs.PrimaryKey
}

func (c *Collection) Put(doc *Document) error {
	docPk, ok := doc.Fields[c.GetPk()]

	if !ok {
		return fmt.Errorf("Error: Document must have field \"key\"")
	}

	isUsedPk := u.Some(c.Documents, func(el *Document, _ int) bool {
		return doc.Fields[c.GetPk()].Value == docPk
	})

	if isUsedPk {
		return fmt.Errorf("Error: This primary key used")
	}

	c.Documents = append(c.Documents, doc)

	return nil
}

func (c *Collection) Get(key string) (bool, *Document) {
	docPointer := u.Find(c.Documents, func(doc *Document, _ int) bool {
		return doc.Fields[c.GetPk()].Value == key
	})

	if docPointer != nil {
		return true, *docPointer
	} else {
		return false, nil
	}
}

func (c *Collection) Delete(key string) bool {
	targetDocIdx := u.FindIndex(c.Documents, func(doc *Document, _ int) bool {
		return doc.Fields[c.GetPk()].Value == key
	})

	if targetDocIdx == -1 {
		return false
	}

	u.Delete(c.Documents, targetDocIdx)
	return true
}

func (s *Collection) List() *[]*Document {
	return &s.Documents
}
