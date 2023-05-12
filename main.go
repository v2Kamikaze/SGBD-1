package main

import (
	"fmt"
	"os"
	"sgbd-1/src/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	/*
		storage := storage.GetStorage()

		fmt.Println(storage.Scan())

		for i := 0; i < 19; i++ {
			fmt.Println(storage.Insert([]byte("ABC")))
		}
		storage.Insert([]byte("AB"))
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
