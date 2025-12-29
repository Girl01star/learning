package users

import (
	"errors"

	"github.com/Girl01star/learning/documentstore"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll *documentstore.Collection
}

func NewService(coll *documentstore.Collection) *Service {
	return &Service{coll: coll}
}

func (s *Service) CreateUser(id, name string) (*User, error) {
	u := &User{ID: id, Name: name}

	doc, err := documentstore.MarshalDocument(u)
	if err != nil {
		return nil, err
	}

	if err = s.coll.Put(*doc); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) ListUsers() ([]User, error) {
	docs := s.coll.List()

	res := make([]User, 0, len(docs))
	for i := range docs {
		var u User
		if err := documentstore.UnmarshalDocument(&docs[i], &u); err != nil {
			return nil, err
		}
		res = append(res, u)
	}

	return res, nil
}

func (s *Service) GetUser(userID string) (*User, error) {
	doc, err := s.coll.Get(userID)
	if err != nil {
		if errors.Is(err, documentstore.ErrDocumentNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	var u User
	if err = documentstore.UnmarshalDocument(doc, &u); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Service) DeleteUser(userID string) error {
	err := s.coll.Delete(userID)
	if err != nil {
		if errors.Is(err, documentstore.ErrDocumentNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}
