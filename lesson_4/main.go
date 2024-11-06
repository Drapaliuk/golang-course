package main

import (
	"fmt"
	ds "lesson_4/documentstore"
)

func main() {
	store := ds.NewStore()
	fmt.Println(store)

	isCreatedCollection, _ := store.CreateCollection(ds.CollectionConfig{
		PrimaryKey: "name",
		Name:       "users",
	})

	fmt.Println("Created collection: ", isCreatedCollection)
	fmt.Println("Collections list: ", store.CollectionsList())

	foundCollection, usersCollection := store.GetCollection("users")
	fmt.Println("Get collection \"users\": ", foundCollection, usersCollection)

	fmt.Println("Users: ", usersCollection.List())

	user := ds.NewDocument(map[string]ds.DocumentField{
		"name": {Type: "string", Value: "Jack"},
		"age":  {Type: "number", Value: 32},
	})

	usersCollection.Put(user)
	fmt.Println("Put: done")

	fmt.Println("Users: ", usersCollection.List())

	putErr := usersCollection.Put(user)
	if putErr != nil {
		fmt.Println("Put error: ", putErr)
	}

	isFound, foundUser := usersCollection.Get("Jack")
	fmt.Println("Get: ", isFound, foundUser)

	isDeleted := usersCollection.Delete("Jack")
	fmt.Println("Delete: ", isDeleted)

	isFoundAfterDelete, foundUserAfterDelete := usersCollection.Get("Jack")
	fmt.Println("Get after delete: ", isFoundAfterDelete, foundUserAfterDelete)
	fmt.Println("Users: ", usersCollection.List())

	store.DeleteCollection("users")
	fmt.Println("Collections list: ", store.CollectionsList())

	foundCollectionAfterDel, usersCollectionAfterDel := store.GetCollection("users")
	fmt.Println("Get collection \"users\" after delete: ", foundCollectionAfterDel, usersCollectionAfterDel)
}
