package documentstore

import (
	"errors"
	u "lesson_7/internal/utils"
)

var ErrDocMustHaveFieldKey = errors.New("document must have field \"key\"")
var ErrDocPrimaryKeyUsed = errors.New("document must have field \"key\"")
var ErrDocumentNotFound = errors.New("document not found")
var ErrDeleteDocument = errors.New("error during document deleting")

type Collection struct {
	Configs   CollectionConfig
	Documents []Document
}

func NewCollection(cfg CollectionConfig) Collection {
	return Collection{
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

func (c *Collection) Put(doc Document) error {
	docPk, ok := doc.Fields[c.GetPk()]

	if !ok {
		return ErrDocMustHaveFieldKey
	}

	isUsedPk := u.Some(c.Documents, func(el Document, _ int) bool {
		return doc.Fields[c.GetPk()].Value == docPk.Value
	})

	if isUsedPk {
		return ErrDocPrimaryKeyUsed
	}

	c.Documents = append(c.Documents, doc)

	return nil
}

func (c *Collection) Get(key string) (*Document, error) {
	docPointer := u.Find(c.Documents, func(doc Document, _ int) bool {
		return doc.Fields[c.GetPk()].Value == key
	})

	if docPointer != nil {
		return docPointer, nil
	} else {
		return nil, ErrDocumentNotFound
	}
}

func (c *Collection) Delete(key string) error {
	targetDocIdx := u.FindIndex(c.Documents, func(doc Document, _ int) bool {
		return doc.Fields[c.GetPk()].Value == key
	})

	if targetDocIdx == -1 {
		return errors.Join(
			ErrDeleteDocument,
			ErrDocumentNotFound,
		)
	}

	c.Documents = u.Delete(c.Documents, targetDocIdx)
	return nil
}

func (s *Collection) List() []Document {
	return s.Documents
}
