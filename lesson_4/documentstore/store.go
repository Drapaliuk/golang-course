package documentstore

type Store struct {
	// ...
}

func NewStore() *Store {
	// TODO: Implement
}

func (s *Store) CreateCollection(cfg *CollectionConfig) (bool, *Collection) {
	// Створюємо нву колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створеня то повертаємо `false` та nil
	// TODO: Implement
}

func (s *Store) GetCollection(name string) (bool, *Collection) {
	// TODO: Implement
}

func (s *Store) DeleteCollection(name string) bool {
	// TODO: Implement
}
