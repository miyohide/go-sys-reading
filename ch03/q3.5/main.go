package main

import (
	"bytes"
	"io"
	"os"
)

func MyCopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	written, err = io.Copy(dst, io.LimitReader(src, n))
	if written == n {
		return n, nil
	}
	if written < n && err == nil {
		err = io.EOF
	}
	return
}

func main() {
	reader := bytes.NewReader([]byte("abcdefghijklmn"))

	MyCopyN(os.Stdout, reader, 5)
}
