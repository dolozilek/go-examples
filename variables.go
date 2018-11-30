package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (person Person) sayHello() {
	fmt.Printf("Hello %s %s\n", person.firstName, person.lastName)
}

func (person Person) copy() Person {
	person.firstName = "This is a copy"
	return person
}

func variables() {
	fmt.Println("Hello world")

	var message = "hello"
	message2 := "world"
	message3 := &message2

	fmt.Println(message, message2, *message3)
	fmt.Println("---")

	pointer(message3)
	fmt.Println(*message3)
	fmt.Println(message2)
	fmt.Println("---")

	value(message2)
	fmt.Println("---")
	fmt.Println(message3)
	fmt.Println(*message3)
	fmt.Println(message2)
	fmt.Println("---")

	person1 := Person{"Martin", "Dolozilek"}
	person2 := Person{"Martin", "Mach"}
	person3 := &Person{"Radim", "Loskot"}

	fmt.Println(person1)
	fmt.Println(&person2)
	fmt.Println(*person3)
	person3.firstName = "Ondrej"
	fmt.Println(person3)

	person1.sayHello()

	fmt.Println("---")
	fmt.Println(person1)
	person4 := person1.copy()
	fmt.Println(&person1)
	fmt.Println(&person4)

}

func value(message string) {
	message = "value"
	fmt.Println(&message)
}

func pointer(message *string) {
	*message = "pointer"
}
