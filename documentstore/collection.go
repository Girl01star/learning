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

func (c *Collection) Put(doc Document) {
	if doc.Fields == nil {
		return
	}

	f, ok := doc.Fields[c.cfg.PrimaryKey]
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

	c.docs[key] = doc
}

func (c *Collection) Get(key string) (*Document, bool) {
	doc, ok := c.docs[key]
	if !ok {
		return nil, false
	}
	copyDoc := doc
	return &copyDoc, true
}

func (c *Collection) Delete(key string) bool {
	if _, ok := c.docs[key]; !ok {
		return false
	}

	delete(c.docs, key)
	return true
}

func (c *Collection) List() []Document {
	res := make([]Document, 0, len(c.docs))
	for _, d := range c.docs {
		res = append(res, d)
	}

	return res
}
