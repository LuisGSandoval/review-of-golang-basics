package main

import (
	"fmt"
	"time"
)

func main4() {
	// <Name> = make(chanel <type>)
	c := make(chan int)

	// Send values into go routines
	go func() {
		c <- 100
	}()

	// Sniff, it's when values are extracted from a channel
	val := <-c
	fmt.Println("El valor es:", val)

	// Send value nº 2 in the same channel
	go func() {
		c <- 101
	}()

	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println("El segundo valor es: ", val)
}
