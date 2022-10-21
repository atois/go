//go:build ignore

package main

import "fmt"

func main() {
	map1 := make(map[string]string)

	map1["name"] = "tukou"
	map1["age"] = "22"
	map1["sex"] = "male"
	map1["address"] = "hello road"

	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["name"] = "lili"
	map2["age"] = "33"
	map2["sex"] = "female"

	fmt.Println(map2)
	//s1 := make(s1[]string)
}
