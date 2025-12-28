package documentstore

type CollectionConfig struct {
	PrimaryKey string
}

type Collection struct {
	cfg  CollectionConfig
	docs map[string]Document

	store *Store
	name  string
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

func (c *Collection) Put(doc Document) error {
	f, ok := doc.Fields[c.cfg.PrimaryKey]
	if !ok || f.Type != DocumentFieldTypeString {
		return ErrInvalidPrimaryKey
	}

	key, ok := f.Value.(string)
	if !ok || key == "" {
		return ErrInvalidPrimaryKey
	}

	_, existed := c.docs[key]
	c.docs[key] = doc

	if c.store != nil {
		if existed {
			c.store.addLog(LogDocumentUpdate, c.name, key)
		} else {
			c.store.addLog(LogDocumentCreate, c.name, key)
		}
	}

	return nil
}

func (c *Collection) Get(key string) (*Document, error) {
	doc, ok := c.docs[key]
	if !ok {
		return nil, ErrDocumentNotFound
	}
	copyDoc := doc
	return &copyDoc, nil
}

func (c *Collection) Delete(key string) error {
	if _, ok := c.docs[key]; !ok {
		return ErrDocumentNotFound
	}

	delete(c.docs, key)

	if c.store != nil {
		c.store.addLog(LogDocumentDelete, c.name, key)
	}

	return nil
}

func (c *Collection) List() []Document {
	res := make([]Document, 0, len(c.docs))
	for _, d := range c.docs {
		res = append(res, d)
	}
	return res
}
