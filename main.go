package main

import (
	"fmt"
	"libs/lib"
)

func main() {
	exept, text, counter := lib.ReadTextFile("file.tsv")
	if exept != 0 {
		fmt.Println("err")
	}

	fmt.Printf("\nResult is:\ncounter: %d\ntext: %s\n", counter, text)
}
