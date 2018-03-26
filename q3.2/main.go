package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	out, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	if _, err := io.CopyN(out, rand.Reader, 1024); err != nil {
		log.Fatal(err)
	}
}
