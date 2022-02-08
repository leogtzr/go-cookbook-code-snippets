package main

import (
	"log"
	"os"
)

func main() {
	data := []byte("Hello World\n")
	err := os.WriteFile("data.txt", data, 0644)
	if err != nil {
		log.Fatal("error: cannot write to file", err)
	}

}
