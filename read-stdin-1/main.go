package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// stop it with CTRL+Z in Windows...
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}