package main

import (
	"fmt"
	ds "lesson_3/documentstore"
)

var docKeys = []string{"doc_1", "doc_2"}

func populateDocumentStore(keys []string) {
	for _, key := range keys {
		doc := ds.NewDocument(key)
		err := ds.Put(doc)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	populateDocumentStore(docKeys)
	testDocKey := "doc____3"

	err := ds.Put(ds.NewDocument(testDocKey))
	if err != nil {
		fmt.Println("Put error: ", err)
	}

	err = ds.Put(ds.NewDocument(testDocKey))

	if err != nil {
		fmt.Println("Put error: ", err)
	}

	docs := ds.List()
	fmt.Println("List: ", docs)

	isExistDoc, doc := ds.Get(testDocKey)
	fmt.Println("Get: ", isExistDoc, doc)

	isDeleted := ds.Delete(testDocKey)
	fmt.Println("Delete: ", isDeleted)

	isExistDocAfterDelete, docAfterDelete := ds.Get(testDocKey)
	fmt.Println("Get after delete: ", isExistDocAfterDelete, docAfterDelete)

}
