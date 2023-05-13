package main

import (
	"sgbd-1/src/storage"
	"sgbd-1/src/tui"
)

func main() {

	storage := storage.GetStorage()

	for i := 0; i < 24; i++ {
		storage.Insert([]byte("A"))
	}

	ui := tui.NewTUI()
	ui.Run()

	/*
		storage := storage.GetStorage()

		fmt.Println(storage.Scan())

		for i := 0; i < 19; i++ {
			fmt.Println(storage.Insert([]byte("ABC")))
		}
		storage.Insert([]byte("ABNewTUI
		storage.Insert([]byte("AB"))
		storage.Insert([]byte("AB"))

		fmt.Println(storage.Insert([]byte("ABCDE")))

		for i := 0; i < 10; i++ {
			storage.Delete([]byte("ABC"))
		}

		storage.ReadFree()
		storage.ReadUsed()

		fmt.Println(storage.Seek([]byte("ABCDE"))) */

}
