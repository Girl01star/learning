package documentstore

type Collection struct {
	cfg  CollectionConfig
	docs map[string]Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func newCollection(cfg *CollectionConfig) *Collection {
	realCfg := CollectionConfig{PrimaryKey: "id"}
	if cfg != nil && cfg.PrimaryKey != "" {
		realCfg.PrimaryKey = cfg.PrimaryKey
	}

	return &Collection{
		cfg:  realCfg,
		docs: make(map[string]Document),
	}
}

func (s *Collection) Put(doc Document) {
	if doc.Fields == nil {
		return
	}

	f, ok := doc.Fields[s.cfg.PrimaryKey]
	if !ok {
		return
	}

	if f.Type != DocumentFieldTypeString {
		return
	}

	key, ok := f.Value.(string)
	if !ok || key == "" {
		return
	}

	s.docs[key] = doc
}

func (s *Collection) Get(key string) (*Document, bool) {
	doc, ok := s.docs[key]
	if !ok {
		return nil, false
	}
	copyDoc := doc
	return &copyDoc, true
}

func (s *Collection) Delete(key string) bool {
	if _, ok := s.docs[key]; !ok {
		return false
	}
	delete(s.docs, key)
	return true
}

func (s *Collection) List() []Document {
	res := make([]Document, 0, len(s.docs))
	for _, d := range s.docs {
		res = append(res, d)
	}
	return res
}
