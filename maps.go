package main

import (
	"fmt"
	"reflect"
)

func maps() {
	var m1 = make(map[string]string)
	m1["hello"] = "world"
	m1["hi"] = "cs"

	var m2 = map[string]string{
		"hello": "world",
		"hi":    "cs",
	}

	eq := reflect.DeepEqual(m1, m2)
	if eq {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	elem, ok := m1["good day"]
	fmt.Println(elem)
	fmt.Println(ok)

}
