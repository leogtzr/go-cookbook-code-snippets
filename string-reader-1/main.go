package main

import (
	"fmt"
	"io"
	"strings"
)

func foo(r io.Reader) error {
	val, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println(string(val))

	return nil
}

func main() {
	name := "Leonardo"
	reader := strings.NewReader(name)
	builder := strings.Builder{}

	builder.WriteString("alrpvw")

	err := foo(reader)
	if err != nil {
		panic(err)
	}
}
