package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Hello, %s!", "world")
	fmt.Println(buf.String())
}

/*
func myHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]bytes("Hello World"))
}
*/
