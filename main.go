package main

import (
	"fmt"
	"sgbd-1/src/storage"
)

func main() {
	storage := storage.GetStorage()

	fmt.Println(storage.Scan())

	for i := 0; i < 10; i++ {
		storage.Insert([]byte("ABC"))
	}
	storage.Insert([]byte("AB"))
	storage.Insert([]byte("AB"))
	storage.Insert([]byte("AB"))

	storage.Insert([]byte("ABCDE"))

	for i := 0; i < 10; i++ {
		storage.Delete([]byte("ABC"))
	}

	storage.Delete([]byte("ABCDE"))

	storage.Delete([]byte("AB"))
	storage.Delete([]byte("AB"))
	storage.Delete([]byte("AB"))

	storage.ReadFree()
	storage.ReadUsed()

}
