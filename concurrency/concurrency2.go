package main

import (
	"fmt"
	"sync"
)

func main() {

	// Waitin groups were created to ensure go routins are fully executed.
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// This is a go routin on a annonymous function also called as a lambda function
	go func() {
		fmt.Println("Heeelllllo")
		wg.Done()
	}()

	go func() {
		fmt.Println("Woooorld")
		wg.Done()
	}()

	fmt.Println("Fin")

	wg.Wait()

	fmt.Println("Exit")

}
