package main

import (
	"fmt"
	"os"
)

func main() {

	for index, element := range os.Args[1:] {
		fmt.Println(index, "=>", element)
	}
}
