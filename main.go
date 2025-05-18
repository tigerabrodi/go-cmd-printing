package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide some args that we echo back for you")
		return
	}



	fmt.Println("Echoing back to you...")

	for _, arg := range args {
		fmt.Println(arg)
	}

}