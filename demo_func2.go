package main

func main() {
	var a = 1
	var b = &a

	*b = 80
	println(a)
	println(b)
}
