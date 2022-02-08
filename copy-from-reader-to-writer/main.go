package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

// using a random 1MB test file
var url string = "http://speedtest.ftp.otenet.gr/files/test1Mb.db"

func readWrite() {
	r, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()
	data, _ := io.ReadAll(r.Body)
	os.WriteFile("rw.data", data, 0755)
}

func copy() {
	r, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()
	file, _ := os.Create("copy.data")
	defer file.Close()
	writer := bufio.NewWriter(file)
	io.Copy(writer, r.Body)
	writer.Flush()
}

func main() {

}
