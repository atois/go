//go:build ignore

package main

import "fmt"

func main() {
	s1 := Student{name: "zhangsan", age: 11}
	fmt.Println(s1.name)

	func() {
		fmt.Println("noname")
	}()

	s2 := struct {
		name string
		age  int
	}{
		name: "lisi",
		age:  12,
	}
	fmt.Println(s2.name, s2.age)

	w2 := Worker{"wangwu", 53}
	fmt.Printf("w2: %v\n", w2)
}

type Student struct {
	name string
	age  int
}

type Worker struct {
	//name string
	//age  int
	string
	int
}
