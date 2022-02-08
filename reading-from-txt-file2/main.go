package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("error: cannot open file: ", err)
	}
	defer file.Close()

	// to get a FileInfo struct.
	stat, err := file.Stat()
	if err != nil {
		log.Fatal("error: cannot read file stats", err)
	}

	data := make([]byte, stat.Size())

	bytes, err := file.Read(data)
	if err != nil {
		log.Fatal("error: cannot read file: ", err)
	}

	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(data))

}
