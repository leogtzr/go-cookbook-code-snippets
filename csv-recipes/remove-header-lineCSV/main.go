package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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
	file, err := os.Open("input.csv")
	if err != nil {
		log.Fatal("error: cannot open CSV file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read() // skip the first line...
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error: cannot read CSV file", err)
	}

	var users []User
	for _, row := range rows {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		user := User{ID: int(id),
			firstName: row[1],
			lastName:  row[2],
			email:     row[3],
		}

		users = append(users, user)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
