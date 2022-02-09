package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := []byte("Hello World!\n")
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal("error: cannot create file", err)
	}
	defer file.Close()

	bytes, err := file.Write(data)
	if err != nil {
		log.Fatal("error: cannot write to file", err)
	}
	fmt.Printf("Wrote %d bytes to file\n", bytes)
}
