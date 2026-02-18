package proto

import (
	"github.com/Girl01star/learning/lesson_12/internal/documentstore"
)

type Request struct {
	Cmd        string `json:"cmd"`
	Collection string `json:"collection,omitempty"`
	Key        string `json:"key,omitempty"`

	PrimaryKey string `json:"primary_key,omitempty"` // для create_collection

	Doc documentstore.Document `json:"doc,omitempty"` // для put
}

type Response struct {
	Ok    bool        `json:"ok"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
