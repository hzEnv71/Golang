package main

import (
	. "fmt"

	"main_test"
	"sort"
) //无需通过包名

func main() {

	main_test.V1()
	a := make([]int, 5)
	a[0] = 3
	a[1] = 1
	a[2] = 2
	var b []int
	b = a[0:2]
	Println(b)
	sort.Ints(b)
	Println(b)
	b = a[1:3]
	Println(b)
}
