package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		[]string{"りんご", "Apple"},
		[]string{"みかん", "Orange"},
		[]string{"スイカ", "Watermelon"},
	}

	w := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			panic(err)
		}
		w.Flush()
	}
}
