package main

import (
	"fmt"

	"github.com/Girl01star/learning/documentstore"
)

func main() {
	store := documentstore.NewStore()

	created, users := store.CreateCollection("users", &documentstore.CollectionConfig{
		PrimaryKey: "key",
	})
	fmt.Println("created:", created)

	doc1 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key":  {Type: documentstore.DocumentFieldTypeString, Value: "user_1"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Alina"},
		},
	}

	doc2 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key":  {Type: documentstore.DocumentFieldTypeString, Value: "user_2"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Bogdan"},
		},
	}

	users.Put(doc1)
	users.Put(doc2)

	fmt.Println("LIST:")
	for _, d := range users.List() {
		fmt.Println("-", d.Fields["key"].Value, d.Fields["name"].Value)
	}

	deleted := users.Delete("user_2")
	fmt.Println("DELETE user_2 ->", deleted)

	_, ok := users.Get("user_2")
	fmt.Println("GET user_2 after delete ->", ok)
}
