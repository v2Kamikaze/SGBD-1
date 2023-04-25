package main

import (
	"fmt"
	"log"
	"sgbd-1/storage"
)

func main() {

	doc, err := storage.NewDocument([]byte("AB"))
	if err != nil {
		log.Fatal(err)
	}

	page := storage.NewPage(0)
	page.AddDocument(doc)
	fmt.Println("Página após inserir o primeiro documento: ", page)
	fmt.Println("Conteúdo da página: ", string(page.Content[:]))

	err = page.AddDocument(doc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Página após inserir o segundo documento:", page)
	fmt.Println("Conteúdo da página: ", string(page.Content[:]))

	err = page.AddDocument(doc)
	if err != nil {
		fmt.Println(err)
	}

}
