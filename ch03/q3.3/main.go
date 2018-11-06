package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	dest, err := os.Create("q3.3.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()

	sReader := strings.NewReader("Readerの出力内容は文字列で渡す")

	writer, err := zipWriter.Create("sReader.txt")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(writer, sReader)
	if err != nil {
		panic(err)
	}
}
