package main

import (
	"fmt"
	"os"
)

func main() {

	// first arg is always the program name
	args := os.Args[1:]

	if len(args) == 1 {
		fmt.Println("Provide some args")
		return
	}

	for _, arg := range args {
		fmt.Println(arg)
	}
}
