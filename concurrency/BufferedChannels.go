package main

import (
	"fmt"
	"time"
)

// Car struct to model the car entity
type Car struct {
	Model string
}

func main5() {
	c := make(chan *Car, 3)

	go func() {
		c <- &Car{Model: "1993"}
		c <- &Car{Model: "1996"}
		c <- &Car{Model: "1965"}
		c <- &Car{Model: "1992"}
		c <- &Car{Model: "2013"}
		c <- &Car{Model: "2016"}
		close(c)
	}()

	for i := range c {
		fmt.Println(i.Model)

		time.Sleep(time.Second * 3)
		fmt.Println("Hola")
	}

}
