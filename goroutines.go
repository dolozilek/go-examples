package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
	wg.Done()
}

func gorountines() {
	wg.Add(3)
	go sayHello("Martin")
	go sayHello("Ondra")
	sayHello("Radim")
	wg.Wait()
}
