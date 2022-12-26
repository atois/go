//go:build ignore

package main

import "fmt"

func main() {
	//creat map
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"go": 95, "to": 27, "bed": 14, "now": 78}
	fmt.Printf("map1: %v\n", map1)
	fmt.Printf("map2: %v\n", map2)
	fmt.Printf("map3: %v\n", map3)

	fmt.Println("-------defult value------------")
	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = ">_<"
	map1[4] = "cod"
	map1[5] = "got"
	map1[6] = "flash"

	//according key to pring value
	fmt.Println(map1)
	fmt.Println(map1[4])
	fmt.Println(map1[40]) //wrong number reture to nil

	v1, ok := map1[40]
	if ok {
		fmt.Printf("v1: %v\n", v1)
	} else {
		fmt.Println("not exist", v1)
	}

	//change value
	fmt.Println(map1)
	map1[3] = "shin"
	fmt.Printf("map1: %v\n", map1)

	//delete value
	delete(map1, 3)
	fmt.Print(map1)
	delete(map1, 30)
	fmt.Println(map1)

}
