package main

import (
	"fmt"
	"log"

	"github.com/Girl01star/learning/documentstore"
)

func main() {
	store := documentstore.NewStore()

	users, err := store.CreateCollection("users", &documentstore.CollectionConfig{
		PrimaryKey: "id",
	})
	if err != nil {
		log.Fatal("CreateCollection:", err)
	}

	err = users.Put(documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id":   {Type: documentstore.DocumentFieldTypeString, Value: "1"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Alina"},
		},
	})
	if err != nil {
		log.Fatal("Put:", err)
	}

	if err := store.DumpToFile("dump.json"); err != nil {
		log.Fatal("DumpToFile:", err)
	}

	newStore, err := documentstore.NewStoreFromFile("dump.json")
	if err != nil {
		log.Fatal("NewStoreFromFile:", err)
	}

	fmt.Println("Logs after restore:")
	for _, l := range newStore.Logs() {
		fmt.Println(l.Action, l.Collection, l.Key)
	}
}
