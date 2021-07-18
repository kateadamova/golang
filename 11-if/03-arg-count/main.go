package main

import (
	"fmt"
	"os"
)

func main() {
	len := len(os.Args) - 1
	if len == 0 {
		fmt.Println("Give me args")
	} else if len == 1 {
		fmt.Printf("There is one: %q", os.Args[1])
	} else if len == 2 {
		fmt.Printf("There are two: %q", os.Args[1]+" "+os.Args[2])
	} else {
		fmt.Printf("There are %d arguments", len)
	}
}
