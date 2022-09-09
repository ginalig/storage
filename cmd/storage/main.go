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

	fmt.Println("Hello world", file)
}
