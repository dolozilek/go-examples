package main

import "fmt"

func arrays() {
	var a [10]int
	a[0] = 5
	a[1] = 10
	fmt.Println(a)

	var slice = a[0:2]
	fmt.Println(slice)

	a[0] = 100
	fmt.Println(slice)

	for index, i := range a {
		fmt.Printf("%v: %v\n", index, i)
	}
}
