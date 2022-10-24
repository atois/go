//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "helloworld"
	fmt.Println(strings.Contains(s1, "abc"))
	fmt.Println(strings.Contains(s1, "llo"))

	fmt.Println(strings.Count(s1, "aaabbbccc"))
	fmt.Println(strings.Count(s1, "o"))

	s2 := "20030303note"
	if strings.HasPrefix(s2, "2003") {
		fmt.Println("long time no see")
	}

	if strings.HasSuffix(s2, "note") {
		fmt.Println("this is note")
	}

	fmt.Println(strings.Index(s2, "0"))
	fmt.Println(strings.IndexAny(s2, "abcdef"))
	fmt.Println(strings.LastIndex(s2, "0"))

	ss1 := []string{"the", "last", "drop", "never"}
	s3 := strings.Join(ss1, "--")
	fmt.Println(s3)

}
