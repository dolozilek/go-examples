package main

import "fmt"

type AdderFunction func(i int) int

func adderBuilder(base int) AdderFunction {
	return func(x int) int {
		base += x
		return base
	}
}

func functions() {
	var f1 AdderFunction = adderBuilder(5)
	var f2 AdderFunction = adderBuilder(10)
	fmt.Println(f1(3))
	fmt.Println(f2(3))
}
