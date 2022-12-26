package main

import "fmt"

func main() {
	age := 16
	r := age >= 18
	fmt.Printf("r: %v\n", r)

	if r {
		fmt.Println("you are the man")
	} else {
		fmt.Println("you are chidren")
	}

}
