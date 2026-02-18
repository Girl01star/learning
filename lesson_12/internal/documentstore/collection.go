package documentstore

import "sync"

type CollectionConfig struct {
	PrimaryKey string
}

type Collection struct {
	mu sync.RWMutex

	cfg  CollectionConfig
	docs map[string]Document

	store *Store
	name  string

	indexes map[string]*Index
}

func newCollection(cfg *CollectionConfig) *Collection {
	realCfg := CollectionConfig{PrimaryKey: "id"}
	if cfg != nil && cfg.PrimaryKey != "" {
		realCfg.PrimaryKey = cfg.PrimaryKey
	}

	return &Collection{
		cfg:     realCfg,
		docs:    make(map[string]Document),
		indexes: make(map[string]*Index),
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

	var existed bool
	var old *Document

	c.mu.Lock()
	{
		if prev, ok := c.docs[key]; ok {
			existed = true
			tmp := prev
			old = &tmp
		}
		c.docs[key] = doc

		c.reindexOnUpsertLocked(key, old, doc)
	}
	c.mu.Unlock()

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
	c.mu.RLock()
	defer c.mu.RUnlock()

	doc, ok := c.docs[key]
	if !ok {
		return nil, ErrDocumentNotFound
	}
	copyDoc := doc
	return &copyDoc, nil
}

func (c *Collection) Delete(key string) error {
	var old Document

	c.mu.Lock()
	{
		doc, ok := c.docs[key]
		if !ok {
			c.mu.Unlock()
			return ErrDocumentNotFound
		}
		old = doc
		delete(c.docs, key)

		c.reindexOnDeleteLocked(key, old)
	}
	c.mu.Unlock()

	if c.store != nil {
		c.store.addLog(LogDocumentDelete, c.name, key)
	}
	return nil
}

func (c *Collection) List() []Document {
	c.mu.RLock()
	defer c.mu.RUnlock()

	res := make([]Document, 0, len(c.docs))
	for _, d := range c.docs {
		res = append(res, d)
	}
	return res
}

type QueryParams struct {
	Desc     bool
	MinValue *string
	MaxValue *string
}

func (c *Collection) CreateIndex(fieldName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.indexes == nil {
		c.indexes = make(map[string]*Index)
	}
	if _, ok := c.indexes[fieldName]; ok {
		return ErrIndexAlreadyExists
	}

	idx := NewIndex(fieldName)

	for key, doc := range c.docs {
		val, ok := getStringField(doc, fieldName)
		if !ok {
			continue
		}
		idx.Add(val, key)
	}

	c.indexes[fieldName] = idx
	return nil
}

func (c *Collection) DeleteIndex(fieldName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.indexes == nil {
		return ErrIndexNotFound
	}
	if _, ok := c.indexes[fieldName]; !ok {
		return ErrIndexNotFound
	}

	delete(c.indexes, fieldName)
	return nil
}

func (c *Collection) Query(fieldName string, params QueryParams) ([]Document, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	idx, ok := c.indexes[fieldName]
	if !ok {
		return nil, ErrIndexNotFound
	}

	keys := idx.RangeKeys(params.MinValue, params.MaxValue, params.Desc)

	res := make([]Document, 0, len(keys))
	for _, key := range keys {
		if doc, ok := c.docs[key]; ok {
			res = append(res, doc)
		}
	}
	return res, nil
}

func getStringField(doc Document, fieldName string) (string, bool) {
	f, ok := doc.Fields[fieldName]
	if !ok || f.Type != DocumentFieldTypeString {
		return "", false
	}
	v, ok := f.Value.(string)
	if !ok || v == "" {
		return "", false
	}
	return v, true
}

func (c *Collection) reindexOnUpsertLocked(key string, oldDoc *Document, newDoc Document) {
	if len(c.indexes) == 0 {
		return
	}

	for fieldName, idx := range c.indexes {
		newVal, newOk := getStringField(newDoc, fieldName)

		if oldDoc == nil {
			if newOk {
				idx.Add(newVal, key)
			}
			continue
		}

		oldVal, oldOk := getStringField(*oldDoc, fieldName)

		switch {
		case oldOk && !newOk:
			idx.Remove(oldVal, key)
		case !oldOk && newOk:
			idx.Add(newVal, key)
		case oldOk && newOk && oldVal != newVal:
			idx.Remove(oldVal, key)
			idx.Add(newVal, key)
		}
	}
}

func (c *Collection) reindexOnDeleteLocked(key string, oldDoc Document) {
	if len(c.indexes) == 0 {
		return
	}
	for fieldName, idx := range c.indexes {
		oldVal, ok := getStringField(oldDoc, fieldName)
		if ok {
			idx.Remove(oldVal, key)
		}
	}
}
