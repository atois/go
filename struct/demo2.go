//go:build ignore

package main

import "fmt"

func main() {
	p1 := Person{"faker", 24, "male", "kor"}
	fmt.Println(p1)
	fmt.Printf("%p %T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p %T\n", &p2, p2)

	p2.name = "xiaohu"
	fmt.Println(p2)
	fmt.Println(p1)

	var pp1 *Person
	pp1 = &p1
	fmt.Printf("pp1: %p %T\n", pp1, pp1)
	fmt.Println(*pp1)

	//(*pp1).name = "yagao"
	pp1.name = "yagao"
	fmt.Println(pp1)
	fmt.Printf("pp1: %p %T\n", pp1, pp1)
	fmt.Println(*pp1)

	pp1.name = "dongdong"
	fmt.Println(pp1)
	fmt.Println(p1)

	pp2 := new(Person)
	fmt.Printf("pp2: %T\n", pp2)
	fmt.Println(pp2)

	pp2.name = "jj"
	pp2.age = 20
	pp2.sex = "male"
	pp2.address = "tokyo"
	fmt.Println(pp2)

	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)
}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
