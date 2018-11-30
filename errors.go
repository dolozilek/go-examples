package main

import (
	"errors"
	"fmt"
)

func getData() (data string, e error) {
	data = "Hello world"
	e = errors.New("error")
	return
}

func errorHandling() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	data, e := getData()
	if e != nil {
		panic("Big problem")
	}

	fmt.Println(data)

}
