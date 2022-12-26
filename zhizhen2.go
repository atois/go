package main

func main() {
	var a = 1
	var b = &a
	*b = 99
	println(a)
	println(b)
	println(*b)
}
