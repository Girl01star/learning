package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/Girl01star/learning/documentstore"
	"github.com/Girl01star/learning/users"
)

func main() {
	store := documentstore.NewStore()
	cfg := documentstore.CollectionConfig{PrimaryKey: "<generated_primary_key>"}
	coll, err := store.CreateCollection("users", &cfg)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create collection"))
		return
	}
	service := users.NewService(coll)
	user, err := service.CreateUser("<genareted_id>", "Bogdan")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create user"))
		return
	}
	fmt.Println(user)
	fmt.Println(service.ListUsers())

}
