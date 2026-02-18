package documentstore

import "errors"

var (
	ErrCollectionAlreadyExists = errors.New("collection already exists")
	ErrCollectionNotFound      = errors.New("collection not found")
	ErrInvalidCollectionName   = errors.New("invalid collection name")

	ErrInvalidPrimaryKey        = errors.New("invalid primary key")
	ErrDocumentNotFound         = errors.New("document not found")
	ErrUnsupportedDocumentField = errors.New("unsupported document field")

	ErrIndexAlreadyExists = errors.New("index already exists")
	ErrIndexNotFound      = errors.New("index not found")
)
