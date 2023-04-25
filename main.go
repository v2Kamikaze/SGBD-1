package main

import (
	"fmt"
	"sgbd-1/storage"
)

func main() {
	page := storage.NewPage(0)
	doc1, _ := storage.NewDocument([]byte("AB"))
	doc2, _ := storage.NewDocument([]byte("CDE"))

	page.AddDocument(doc1)
	fmt.Println(page)

	page.AddDocument(doc2)
	fmt.Println(page)

	/* 	fmt.Println("Está cheia? ", page.IsFull())

	   	page.DeleteDocument(doc1.Content)
	   	fmt.Println(page)
	   	fmt.Println(page.IsEmpty())
	   	fmt.Println("Está cheia? ", page.IsFull())

	   	page.DeleteDocument(doc2.Content)
	   	fmt.Println(page)
	   	fmt.Println(page.IsEmpty())
	   	fmt.Println("Está cheia? ", page.IsFull())

	   	page.DeleteDocument(doc2.Content)
	   	fmt.Println("Está cheia? ", page.IsFull()) */

	fmt.Println(storage.GetStorage())
	fmt.Println(storage.GetStorage())

	did, err := page.GetDID(doc2.Content[:1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(did)

}
