package documentstore

import (
	"errors"
	u "lesson_6/utils"
)

var ErrCollectionAlreadyExists = errors.New("collection already exists")
var ErrCollectionNotFound = errors.New("collection not found")
var ErrDeleteCollection = errors.New("error during document collection")

type Store struct {
	Collections []Collection
}

func NewStore() Store {
	return Store{}
}

func (s *Store) CreateCollection(cfg CollectionConfig) (*Collection, error) {
	isUsedName := u.Some(s.Collections, func(el Collection, _ int) bool {
		return el.Configs.Name == cfg.Name
	})

	if isUsedName {
		return nil, ErrCollectionAlreadyExists
	}

	newCollection := NewCollection(cfg)
	s.Collections = append(s.Collections, newCollection)

	logAction("Created collection", newCollection.Configs.Name)

	return &newCollection, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	collection := u.Find(s.Collections, func(c Collection, _ int) bool {
		return c.Configs.Name == name
	})

	if collection == nil {
		return nil, ErrCollectionNotFound
	} else {
		return collection, nil
	}
}

func (s *Store) DeleteCollection(name string) error {
	collectionIdx := u.FindIndex(s.Collections, func(el Collection, _ int) bool {
		return el.Configs.Name == name
	})

	if collectionIdx == -1 {
		return errors.Join(
			ErrDeleteCollection,
			ErrCollectionNotFound,
		)
	}

	logAction("Deleted collection", name)

	s.Collections = u.Delete(s.Collections, collectionIdx)

	return nil
}

func (s *Store) CollectionsList() []Collection {
	return s.Collections
}
