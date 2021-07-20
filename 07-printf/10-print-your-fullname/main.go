package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Your name is %s and your lastname is %s.", os.Args[1], os.Args[2])
}
