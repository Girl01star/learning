package main

import (
	"fmt"

	"github.com/Girl01star/learning/documentstore"
)

func main() {
	doc1 := &documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key":  {Type: documentstore.DocumentFieldTypeString, Value: "user_1"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Alina"},
			"age":  {Type: documentstore.DocumentFieldTypeNumber, Value: 26},
		},
	}

	doc2 := &documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key":    {Type: documentstore.DocumentFieldTypeString, Value: "user_2"},
			"active": {Type: documentstore.DocumentFieldTypeBool, Value: true},
		},
	}

	documentstore.Put(doc1)
	documentstore.Put(doc2)

	fmt.Println("LIST:")
	for _, d := range documentstore.List() {
		fmt.Println(" - key =", d.Fields["key"].Value)
	}

	deleted := documentstore.Delete("user_2")
	fmt.Println("DELETE user_2 ->", deleted)

	_, ok := documentstore.Get("user_2")
	fmt.Println("GET user_2 after delete ->", ok)
}
