package main

import (
	"fmt"

	"github.com/Girl01star/learning/documentstore"
)

func main() {
	store := documentstore.NewStore()

	users, _ := store.CreateCollection("users", &documentstore.CollectionConfig{
		PrimaryKey: "id",
	})

	_ = users.Put(documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id":   {Type: documentstore.DocumentFieldTypeString, Value: "1"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Alina"},
		},
	})

	_ = store.DumpToFile("dump.json")

	newStore, _ := documentstore.NewStoreFromFile("dump.json")

	fmt.Println("Logs after restore:")
	for _, l := range newStore.Logs() {
		fmt.Println(l.Action, l.Collection, l.Key)
	}
}
