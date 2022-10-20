//go:build ignore

package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("--------------creat----------------")
	s1 := a[:5]  //1-5  qiankaihoubi
	s2 := a[3:8] //4-8
	s3 := a[5:]
	s4 := a[:]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)

	fmt.Println("-----------------len-cap---------------")
	fmt.Printf("s1	len:%d,cap:%d\n", len(s1), cap(s1)) //cap:from strat to end
	fmt.Printf("s2	len:%d,cap:%d\n", len(s2), cap(s2))
	fmt.Printf("s3	len:%d,cap:%d\n", len(s3), cap(s3))
	fmt.Printf("s4	len:%d,cap:%d\n", len(s4), cap(s4))

	fmt.Println("------------------appand---------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)

	fmt.Println("-----------------len-cap---------------")
	fmt.Printf("s1	len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("s2	len:%d,cap:%d\n", len(s2), cap(s2))
	fmt.Printf("s3	len:%d,cap:%d\n", len(s3), cap(s3))
	fmt.Printf("s4	len:%d,cap:%d\n", len(s4), cap(s4))
}
