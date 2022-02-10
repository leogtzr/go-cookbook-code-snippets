package main

import (
	"log"
	"os"
)

func main() {
	tmpdir, err := os.MkdirTemp(os.TempDir(), "mytmpdir_*")
	if err != nil {
		log.Fatal("error: cannot create temp directory", err)
	}
	defer os.RemoveAll(tmpdir)

	tmpfile, err := os.CreateTemp(tmpdir, "mytmp_*")
	if err != nil {
		log.Fatal("error: cannot create temp file", err)
	}

	data := []byte("Some random stuff for the temporary file")
	_, err = tmpfile.Write(data)
	if err != nil {
		log.Fatal("error: cannot write to temp file", err)
	}

	err = tmpfile.Close()
	if err != nil {
		log.Fatal("error: cannot close temp file", err)
	}
}
