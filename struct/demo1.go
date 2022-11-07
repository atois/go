//go:build ignore

package main

import "fmt"

func main() {
	var p1 Person
	fmt.Println(p1)
	p1.name = "zhangsan"
	p1.age = 30
	p1.sex = "male"
	p1.address = "beijing"

	fmt.Printf("name: %s, age: %d, sex: %s, address:%s\n", p1.name, p1.age, p1.sex, p1.address)

	p2 := Person{}
	p2.name = "ruby"
	p2.age = 18
	p2.sex = "female"
	p2.address = "nanjing"
	fmt.Printf("name: %s, age: %d, sex: %s, address:%s\n", p2.name, p2.age, p2.sex, p2.address)

	p3 := Person{name: "zhizi", age: 14, sex: "female", address: "dongjing"}
	fmt.Printf("p3: %v\n", p3)

	p4 := Person{
		name:    "lisi",
		age:     40,
		sex:     "male",
		address: "xian",
	}
	fmt.Printf("p4: %v\n", p4)

	p5 := Person{"jim", 34, "male", "hegang"}
	fmt.Printf("p5: %v\n", p5)

}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
