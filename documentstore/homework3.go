package documentstore

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value any
}

type Document struct {
	Fields map[string]DocumentField
}

var documents = map[string]*Document{}

func Put(doc *Document) {
	if doc == nil || doc.Fields == nil {
		return
	}
	keyField, ok := doc.Fields["key"]
	if !ok {
		return
	}

	if keyField.Type != DocumentFieldTypeString {
		return
	}

	key, ok := keyField.Value.(string)
	if !ok {
		return
	}
	if key == "" {
		return
	}
	documents[key] = doc
}

func Get(key string) (*Document, bool) {
	doc, ok := documents[key]
	if !ok {
		return nil, false
	}
	return doc, true
}

func Delete(key string) bool {
	_, ok := documents[key]
	if !ok {
		return false
	}
	delete(documents, key)
	return true
}

func List() []*Document {
	result := make([]*Document, 0, len(documents))
	for _, doc := range documents {
		result = append(result, doc)
	}
	return result
}
