package documentstore

import (
	"encoding/json"
)

func MarshalDocument(input any) (*Document, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	doc := &Document{Fields: make(map[string]DocumentField, len(m))}
	for k, v := range m {
		f, err := valueToField(v)
		if err != nil {
			return nil, err
		}
		doc.Fields[k] = f
	}

	return doc, nil
}

func UnmarshalDocument(doc *Document, output any) error {
	m := make(map[string]any, len(doc.Fields))
	for k, f := range doc.Fields {
		v, err := fieldToValue(f)
		if err != nil {
			return err
		}
		m[k] = v
	}

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, output)
}

func valueToField(v any) (DocumentField, error) {
	switch x := v.(type) {
	case string:
		return DocumentField{Type: DocumentFieldTypeString, Value: x}, nil

	case bool:
		return DocumentField{Type: DocumentFieldTypeBool, Value: x}, nil

	case float64:
		return DocumentField{Type: DocumentFieldTypeNumber, Value: x}, nil

	case []any:
		// храним массив как []any (без превращения элементов в DocumentField)
		return DocumentField{Type: DocumentFieldTypeArray, Value: x}, nil

	case map[string]any:
		// вложенный объект
		return DocumentField{Type: DocumentFieldTypeObject, Value: x}, nil

	case nil:
		// JSON null — можно хранить как object=nil или unsupported.
		// Сделаем проще: поддержим как object(nil)
		return DocumentField{Type: DocumentFieldTypeObject, Value: nil}, nil

	default:
		return DocumentField{}, ErrUnsupportedDocumentField
	}
}

func fieldToValue(f DocumentField) (any, error) {
	switch f.Type {
	case DocumentFieldTypeString:
		s, ok := f.Value.(string)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		return s, nil

	case DocumentFieldTypeBool:
		b, ok := f.Value.(bool)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		return b, nil

	case DocumentFieldTypeNumber:
		return f.Value, nil

	case DocumentFieldTypeArray:
		a, ok := f.Value.([]any)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		return a, nil

	case DocumentFieldTypeObject:
		if f.Value == nil {
			return nil, nil
		}
		o, ok := f.Value.(map[string]any)
		if !ok {
			return nil, ErrUnsupportedDocumentField
		}
		return o, nil

	default:
		return nil, ErrUnsupportedDocumentField
	}
}
