package main

import "fmt"

func main() {
	s1 := "tom"
	s2 := "20"
	msg := fmt.Sprintf("%s,%s", s1, s2)
	fmt.Printf("msg: %v\n", msg)
}
