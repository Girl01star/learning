package documentstore

type Store struct {
	collections map[string]*Collection
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {

	if name == "" {
		return false, nil
	}

	if _, exists := s.collections[name]; exists {
		return false, nil
	}

	col := newCollection(cfg)
	s.collections[name] = col
	return true, col
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	col, ok := s.collections[name]
	return col, ok
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.collections[name]; !ok {
		return false
	}
	delete(s.collections, name)
	return true
}
