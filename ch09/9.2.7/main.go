package main

import (
	"fmt"
	"os"
)

func main()  {
	info1, err := os.Stat("/Users/miyohide/.config/nvim/init.vim")
	if err != nil {
		panic(err)
	}
	info2, err := os.Stat("/Users/miyohide/work/dotfiles/init.vim")
	if err != nil {
		panic(err)
	}
	if os.SameFile(info1, info2) {
		fmt.Println("同じファイル")
	} else {
		fmt.Println("違うファイル")
	}
}
