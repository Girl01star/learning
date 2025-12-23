package documentstore

import "time"

func now() time.Time { return time.Now() }

type Store struct {
	collections map[string]*Collection
	logs        []LogEntry
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
		logs:        make([]LogEntry, 0),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if name == "" {
		return nil, ErrCollectionNotFound
	}
	if _, exists := s.collections[name]; exists {
		return nil, ErrCollectionAlreadyExists
	}

	col := newCollection(cfg, s, name)
	s.collections[name] = col

	s.addLog(LogEntry{
		At:         now(),
		Action:     LogCollectionCreate,
		Collection: name,
	})
	return col, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	col, ok := s.collections[name]
	if !ok {
		return nil, ErrCollectionNotFound
	}
	return col, nil
}

func (s *Store) DeleteCollection(name string) error {
	if _, ok := s.collections[name]; !ok {
		return ErrCollectionNotFound
	}
	delete(s.collections, name)

	s.addLog(LogEntry{
		At:         now(),
		Action:     LogCollectionDelete,
		Collection: name,
	})
	return nil
}
