package documentstore

import "errors"

var (
	ErrDocumentNotFound         = errors.New("document not found")
	ErrCollectionAlreadyExists  = errors.New("collection already exists")
	ErrCollectionNotFound       = errors.New("collection not found")
	ErrInvalidCollectionName    = errors.New("invalid collection name")
	ErrInvalidPrimaryKey        = errors.New("invalid primary key")
	ErrUnsupportedDocumentField = errors.New("unsupported document field")
)
