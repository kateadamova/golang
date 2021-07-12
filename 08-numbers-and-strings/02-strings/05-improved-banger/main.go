package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg, length := os.Args[1], utf8.RuneCountInString(os.Args[1])
	s := msg + strings.Repeat("!", length)
	fmt.Println(s)
}
