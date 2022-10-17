package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	site := WebSite{Name: "duoke360"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)
	fmt.Printf("site: %T\n", site)

	a := 100
	fmt.Printf("a: %T\n", a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %b\n", a)
	fmt.Printf("a: %c\n", a)

	b := true
	fmt.Printf("b: %t\n", b)

	x := 100
	p := &x
	fmt.Printf("p: %p\n", p)

}
