//go:build ignore

package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	map2 := make(map[string]float64)
	fmt.Printf("map1: %T\n", map1)
	fmt.Printf("map1: %T\n", map2)

	map3 := make(map[string]map[string]string)
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "30"
	m1["salary"] = "3000"

	map3["hr"] = m1

	m2 := make(map[string]string)
	m2["name"] = "ruby"
	m2["age"] = "28"
	m2["salary"] = "8000"

	map3["manger"] = m2
	fmt.Println(map3)

	fmt.Println("a---------------------------")

	map4 := make(map[string]string)
	map4["kongkong"] = "the last dorp"
	map4["powder"] = "yeluoli"
	map4["weiwei"] = "big sister"

	fmt.Println(map4)

	map5 := map4
	fmt.Println(map5)

	map5["kongkong"] = "the frist drop!"

	fmt.Println(map4)
	fmt.Println(map5)
}
