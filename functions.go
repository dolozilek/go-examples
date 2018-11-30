package main

import "fmt"

type HelloFunction func(string) string

func (f HelloFunction) fromPet(pet Pet) HelloFunction {
	return func(s string) string {
		return f(pet.name)
	}
}

type Pet struct {
	name string
}

func main() {
	var f HelloFunction = func(name string) string {
		return fmt.Sprintf("Hello %s\n", name)
	}

	pet := Pet{"bobika"}
	fmt.Println(f.fromPet(pet)("hello"))
}
