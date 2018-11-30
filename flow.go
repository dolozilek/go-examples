package main

import (
	"fmt"
	"math"
)

func flow() {
	hello := "hello"

	if len(hello) > 2 {
		fmt.Printf("Hello has: %v characters\n", len(hello))
	}

	switch hello {
	case "hello":
		fmt.Println("it's hello")
	case "hi":
		fmt.Println("hi")
	default:
		panic("unknown type")
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//while
	sum = 0
	for sum < 1000 {
		sum += 5
	}
	fmt.Println(sum)

	//init statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
