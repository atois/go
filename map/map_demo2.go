//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]string)
	map1[1] = "hong kong"
	map1[2] = "beijing"
	map1[3] = "xian"
	map1[4] = "shanghai"

	//traverse
	for a, b := range map1 {
		fmt.Println(a, b)
	}

	fmt.Println("------------------")

	for i := 1; i <= len(map1); i++ {
		fmt.Printf("map1[i]: %v\n", map1[i])
	}

	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	//import sort package to sorting
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	s1 := []string{"hard", "easy", "shin", "123", "985", "@#@#$"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)

}
