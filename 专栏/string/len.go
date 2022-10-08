package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "HEllo WorlD"
	fmt.Printf("len(s) : %v\n", len(s))
	fmt.Printf("strings.Split(s,\" \"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))

	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))

	fmt.Printf("strings.HasPrefix(\"HEllo\"): %v\n", strings.HasPrefix(s, "HEllo"))
	fmt.Printf("strings.HasSuffix(\"WorlD\"): %v\n", strings.HasSuffix(s, "WorlD"))

	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll"))
}
