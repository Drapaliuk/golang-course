package documentstore

import u "lesson_3/utils"

type Store struct {
	Collections []Collection
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) CreateCollection(cfg *CollectionConfig) (bool, *Collection) {
	isUsedName := u.Some(s.Collections, func(el Collection, _ int) bool {
		return el.Configs.Name == cfg.Name
	})

	if isUsedName {
		return false, nil
	}

	return true, NewCollection(*cfg)
}

func (s *Store) GetCollection(name string) (bool, *Collection) {
	collection := u.Find(s.Collections, func(el Collection, _ int) bool {
		return el.Configs.Name == name
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

	// Зробити щоб delete повертав true/false
	s.Collections = u.Delete(s.Collections, collectionIdx)
}
