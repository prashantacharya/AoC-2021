package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	data, _ := os.ReadFile(fileName)

	fmt.Println(string(data))
}
