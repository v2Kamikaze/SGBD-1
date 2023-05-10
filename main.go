package main

import (
	"fmt"
	"sgbd-1/src/list"
	"sgbd-1/src/page"
)

func main() {
	/* page := page.New(0)
	doc1, _ := doc.New([]byte("AB"))
	doc2, _ := doc.New([]byte("CDE"))

	page.AddDocument(doc1)
	fmt.Println(page)

	page.AddDocument(doc2)
	fmt.Println(page)

	fmt.Println("Está cheia? ", page.IsFull())

	page.DeleteDocument(doc1.Content)
	fmt.Println(page)
	fmt.Println(page.IsEmpty())
	fmt.Println("Está cheia? ", page.IsFull())

	page.DeleteDocument(doc2.Content)
	fmt.Println(page)
	fmt.Println(page.IsEmpty())
	fmt.Println("Está cheia? ", page.IsFull())

	page.DeleteDocument(doc2.Content)
	fmt.Println("Está cheia? ", page.IsFull())

	if did, err := page.GetDID(doc2.Content[:]); err != nil {
		fmt.Println(did)
	} */

	list := list.NewPageList()

	for i := 0; i < 10; i++ {
		list.Add(page.New(i))
	}

	list.ReadAll()

	for i := 0; i < 11; i++ {

		p, err := list.DeletePage(i)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("Página removida => ", p)

	}

	list.ReadAll()

}
