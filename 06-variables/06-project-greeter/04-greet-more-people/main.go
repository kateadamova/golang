package main

import (
	"fmt"
	"os"
)

func main() {
	amount := len(os.Args) - 1
	arg1, arg2, arg3 := os.Args[1], os.Args[2], os.Args[3]
	fmt.Println("There are", amount, "people !")
	fmt.Println("Hello great", arg1, " !")
	fmt.Println("Hello great", arg2, " !")
	fmt.Println("Hello great", arg3, " !")
	fmt.Println("Nice to meet you all.")
}
