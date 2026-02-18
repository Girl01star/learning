package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/Girl01star/learning/documentstore"
)

func main() {
	store := documentstore.NewStore()

	col, err := store.CreateCollection("users", &documentstore.CollectionConfig{PrimaryKey: "id"})
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			defer wg.Done()

			id := strconv.Itoa(i)

			_ = col.Put(documentstore.Document{
				Fields: map[string]documentstore.DocumentField{
					"id":   {Type: documentstore.DocumentFieldTypeString, Value: id},
					"name": {Type: documentstore.DocumentFieldTypeString, Value: "User_" + id},
				},
			})

			_, _ = col.Get(id)

			// иногда удаляем
			if i%3 == 0 {
				_ = col.Delete(id)
			}
		}()
	}

	wg.Wait()
	fmt.Println("done; logs =", len(store.Logs()))
}
