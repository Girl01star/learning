package documentstore

import (
	"encoding/json"
	"os"
)

type dumpCollection struct {
	Cfg  CollectionConfig    `json:"cfg"`
	Docs map[string]Document `json:"docs"`
}

type dumpStore struct {
	Collections map[string]dumpCollection `json:"collections"`
	Logs        []LogEntry                `json:"logs"`
}

func (s *Store) Dump() ([]byte, error) {
	ds := dumpStore{
		Collections: make(map[string]dumpCollection),
		Logs:        s.Logs(),
	}

	for name, col := range s.collections {
		docsCopy := make(map[string]Document)
		for k, v := range col.docs {
			docsCopy[k] = v
		}

		ds.Collections[name] = dumpCollection{
			Cfg:  col.cfg,
			Docs: docsCopy,
		}
	}

	return json.Marshal(ds)
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	var ds dumpStore
	if err := json.Unmarshal(dump, &ds); err != nil {
		return nil, err
	}

	st := NewStore()
	st.logs = append(st.logs, ds.Logs...)

	for name, dc := range ds.Collections {
		col := newCollection(&dc.Cfg)
		col.store = st
		col.name = name

		for k, v := range dc.Docs {
			col.docs[k] = v
		}

		st.collections[name] = col
	}

	return st, nil
}

func NewStoreFromFile(filename string) (*Store, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewStoreFromDump(b)
}

func (s *Store) DumpToFile(filename string) error {
	b, err := s.Dump()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, b, 0644)
}
