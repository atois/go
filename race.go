package main

import "fmt"

func main() {
	var second float32
	var gender string
	fmt.Printf("赵馨语，说出你所用的时间")
	fmt.Scanln(&second)
	fmt.Printf("赵馨语，你是男孩，还是女孩")
	fmt.Scanln(&gender)
	if second <= 8 {
		if gender == "man" {
			fmt.Printf("man is man")
		} else {
			fmt.Printf("girl is girl")
		}
	}

}
