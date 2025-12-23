package main

import (
	"fmt"

	"github.com/Girl01star/learning/documentstore"
)

func main() {
	st := documentstore.NewStore()

	users, _ := st.CreateCollection("users", &documentstore.CollectionConfig{PrimaryKey: "id"})

	_ = users.Put(documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id":   {Type: documentstore.DocumentFieldTypeString, Value: "u1"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Alina"},
		},
	})

	_ = users.Put(documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id":   {Type: documentstore.DocumentFieldTypeString, Value: "u1"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Yura"}, // update
		},
	})

	_ = users.Delete("u1")

	// Dump в файл
	_ = st.DumpToFile("dump.json")

	// Load из файла
	st2, _ := documentstore.NewStoreFromFile("dump.json")

	fmt.Println("Logs loaded:")
	for _, l := range st2.Logs() {
		fmt.Println(l.Action, l.Collection, l.Key, l.At)
	}
}
