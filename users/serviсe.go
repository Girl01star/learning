package users

import (
	"errors"

	"github.com/Girl01star/learning/documentstore"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Collection interface {
	Put(doc documentstore.Document) error
	Get(key string) (*documentstore.Document, error)
	Delete(key string) error
	List() []documentstore.Document
}

type Service struct {
	coll Collection
}

func NewService(coll Collection) *Service {
	return &Service{coll: coll}
}

func (s *Service) Create(u User) error {
	if u.ID == "" {
		return errors.New("empty id")
	}

	doc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id":   {Type: documentstore.DocumentFieldTypeString, Value: u.ID},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: u.Name},
		},
	}
	return s.coll.Put(doc)
}

func (s *Service) Get(id string) (*User, error) {
	doc, err := s.coll.Get(id)
	if err != nil {
		return nil, err
	}

	idField, ok := doc.Fields["id"]
	if !ok {
		return nil, documentstore.ErrUnsupportedDocumentField
	}
	nameField, ok := doc.Fields["name"]
	if !ok {
		return nil, documentstore.ErrUnsupportedDocumentField
	}

	idStr, ok := idField.Value.(string)
	if !ok {
		return nil, documentstore.ErrUnsupportedDocumentField
	}
	nameStr, ok := nameField.Value.(string)
	if !ok {
		return nil, documentstore.ErrUnsupportedDocumentField
	}

	return &User{ID: idStr, Name: nameStr}, nil
}

func (s *Service) Delete(id string) error {
	return s.coll.Delete(id)
}

func (s *Service) List() ([]User, error) {
	docs := s.coll.List()
	res := make([]User, 0, len(docs))

	for _, d := range docs {
		idField, ok := d.Fields["id"]
		if !ok {
			return nil, documentstore.ErrUnsupportedDocumentField
		}
		nameField, ok := d.Fields["name"]
		if !ok {
			return nil, documentstore.ErrUnsupportedDocumentField
		}

		idStr, ok := idField.Value.(string)
		if !ok {
			return nil, documentstore.ErrUnsupportedDocumentField
		}
		nameStr, ok := nameField.Value.(string)
		if !ok {
			return nil, documentstore.ErrUnsupportedDocumentField
		}

		res = append(res, User{ID: idStr, Name: nameStr})
	}

	return res, nil
}
