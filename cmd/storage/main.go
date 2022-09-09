package main

import (
	"fmt"

	"github.com/ginalig/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello world", st)
}
