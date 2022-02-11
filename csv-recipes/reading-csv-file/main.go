package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		log.Fatal("error: cannot open CSV file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error: cannot read CSV file", err)
	}

	for _, row := range rows {
		fmt.Println(row[0])
	}
}
