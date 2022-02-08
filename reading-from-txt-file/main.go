package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Println("error: cannot read file: ", err)
	}

	fmt.Println(string(data))
}
