package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "numeric = %d, string = %s, float = %f\n", 10, "hoge", 12.3)
}
