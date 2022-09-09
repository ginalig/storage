package main

import (
	"fmt"
	"log"

	"github.com/ginalig/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello world", restoredFile)
}
