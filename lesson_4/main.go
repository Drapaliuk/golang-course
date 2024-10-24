package main

import (
	"fmt"
	ds "lesson_3/documentstore"
)

var docKeys = []string{"doc_1", "doc_2", "doc_3", "doc_4", "doc_5"}

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
}
