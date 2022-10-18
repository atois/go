package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "tom"
	age := "20"
	s := strings.Join([]string{name, age}, ",")
	fmt.Printf("s:%v\n", s)

}
