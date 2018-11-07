package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// UNIXのwhichコマンドをGo言語で実装したもの
func main()  {
	// 引数がない場合は使い方を示して終了する
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	// 環境変数PATHを分解して、指定されたファイルが有るかどうかを確認する
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		execpath := filepath.Join(path, os.Args[1])
		_, err := os.Stat(execpath)
		if !os.IsNotExist(err) {
			fmt.Println(execpath)
			return
		}
	}
	os.Exit(1)
}
