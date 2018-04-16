package main

import (
	"fmt"
	"time"
)

func main() {
	w := make(chan string)

	select {
	case <-w:
	case <-time.After(10 * time.Second):
		fmt.Println("10秒経過")
	}
}
