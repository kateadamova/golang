package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}
	fe := os.Args[1]
	if fe == "a" || fe == "e" || fe == "i" || fe == "o" || fe == "u" {
		fmt.Printf("%q is a vowel.", fe)
	} else if fe == "w" || fe == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.", fe)
	} else {
		fmt.Printf("%q is a consonant", fe)
	}
}
