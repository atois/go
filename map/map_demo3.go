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

	map3 := map[string]string{"name": "zhangshan", "age": "28", "sex": "male", "adress": "far away"}
	fmt.Println(map3)

	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)

	for i, val := range s1 {
		fmt.Printf("%d information : \n", i+1)
		fmt.Printf("\t name:%s\n", val["name"])
		fmt.Printf("\t age:%s\n", val["age"])
		fmt.Printf("\t sex:%s\n", val["sex"])
		fmt.Printf("\t address:%s\n", val["address"])

	}

}
