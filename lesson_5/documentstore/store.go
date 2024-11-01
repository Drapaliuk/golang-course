package documentstore

import (
	u "lesson_5/utils"
)

type Store struct {
	Collections []Collection
}

func NewStore() Store {
	return Store{}
}

func (s *Store) CreateCollection(cfg CollectionConfig) (bool, *Collection) {
	isUsedName := u.Some(s.Collections, func(el Collection, _ int) bool {
		return el.Configs.Name == cfg.Name
	})

	if isUsedName {
		return false, nil
	}

	newCollection := NewCollection(cfg)
	s.Collections = append(s.Collections, newCollection)

	return true, &newCollection
}

func (s *Store) GetCollection(name string) (bool, *Collection) {
	collection := u.Find(s.Collections, func(c Collection, _ int) bool {
		return c.Configs.Name == name
	})

	if collection == nil {
		return false, nil
	} else {
		return true, collection
	}
}

func (s *Store) DeleteCollection(name string) bool {
	collectionIdx := u.FindIndex(s.Collections, func(el Collection, _ int) bool {
		return el.Configs.Name == name
	})

	if collectionIdx == -1 {
		return false
	}

	s.Collections = u.Delete(s.Collections, collectionIdx)

	return true
}

func (s *Store) CollectionsList() []Collection {
	return s.Collections
}
