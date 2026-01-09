package documentstore

import "log/slog"

type Store struct {
	collections map[string]*Collection
	logs        []LogEntry

	logger *slog.Logger
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
		logs:        make([]LogEntry, 0),
		logger:      slog.Default(),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if name == "" {
		return nil, ErrInvalidCollectionName
	}

	if _, exists := s.collections[name]; exists {
		return nil, ErrCollectionAlreadyExists
	}

	col := newCollection(cfg)
	col.store = s
	col.name = name

	s.collections[name] = col
	s.addLog(LogCollectionCreate, name, "")

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
	s.addLog(LogCollectionDelete, name, "")

	return nil
}
func (s *Store) Logs() []LogEntry {

	res := make([]LogEntry, len(s.logs))
	copy(res, s.logs)
	return res
}
