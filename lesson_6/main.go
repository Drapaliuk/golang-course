package main

import (
	"fmt"
	"lesson_6/documentstore"
	"log"
)

func main() {
	store := documentstore.NewStore()

	collectionCfg := documentstore.CollectionConfig{
		PrimaryKey: "ID",
		Name:       "users",
	}
	collection, err := store.CreateCollection(collectionCfg)
	if err != nil {
		log.Fatal("Error creating collection:", err)
	}

	fmt.Println("Collection created:", collection.Configs.Name)

	document := map[string]documentstore.DocumentField{
		"ID":   {Type: documentstore.DocumentFieldTypeString, Value: "1"},
		"Name": {Type: documentstore.DocumentFieldTypeString, Value: "Jack"},
	}

	err = collection.Put(documentstore.NewDocument(document))
	if err != nil {
		log.Fatal("Error adding document:", err)
	}

	fmt.Println("Document added to collection")
}
