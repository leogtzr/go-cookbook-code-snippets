package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// User ...
type User struct {
	ID        int
	firstName string
	lastName  string
	email     string
}

func (u User) String() string {
	return fmt.Sprintf("%d - %s %s %s", u.ID, u.firstName, u.lastName, u.email)
}

func main() {
	file, err := os.Create("new_user.csv")
	if err != nil {
		log.Fatal("error: cannot create csv file", err)
	}
	defer file.Close()

	data := [][]string{
		{"id", "first_name", "last_name", "email"},
		{"1", "Sausheong", "Chang", "sausheong@email.com"},
		{"2", "John", "Doe", "john@email.com"},
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal("error: cannot write to CSV file", err)
	}
}
