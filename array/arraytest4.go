package main

import "fmt"

func main() {
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("address: %p\n", &a2)
	fmt.Printf("a2: %d\n", len(a2))
	fmt.Printf("a2: %d\n", len(a2[0]))

	for i := 0; i < len(a2); i++ {
		for j := 0; j < len(a2[i]); j++ {
			fmt.Printf("a2[i][j]: %v\n", a2[i][j])
		}
		fmt.Printf("a2[i]: %v\n", a2[i])
	}
	fmt.Println("----------------------------")
	for _, arr := range a2 {
		for _, val := range arr {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}

}
