package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := strings.TrimRight("inanç           ", " ")
	l := utf8.RuneCountInString(name)
	fmt.Println(l)
}
