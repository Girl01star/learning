package documentstore

import "sync"

type Store struct {
	mu sync.RWMutex

	collections map[string]*Collection
	logs        []LogEntry
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
		logs:        make([]LogEntry, 0),
	}
}
func (s *Store) GetCollection(name string) (*Collection, error) {
	return s.Collection(name)
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if name == "" {
		return nil, ErrInvalidCollectionName
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.collections[name]; ok {
		return nil, ErrCollectionAlreadyExists
	}

	col := newCollection(cfg)
	col.store = s
	col.name = name
	s.collections[name] = col

	s.addLogNoLock(LogCollectionCreate, name, "")
	return col, nil
}

func (s *Store) Collection(name string) (*Collection, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	col, ok := s.collections[name]
	if !ok {
		return nil, ErrCollectionNotFound
	}
	return col, nil
}

func (s *Store) DeleteCollection(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.collections[name]; !ok {
		return ErrCollectionNotFound
	}
	delete(s.collections, name)

	s.addLogNoLock(LogCollectionDelete, name, "")
	return nil
}

func (s *Store) Logs() []LogEntry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cp := make([]LogEntry, len(s.logs))
	copy(cp, s.logs)
	return cp
}
