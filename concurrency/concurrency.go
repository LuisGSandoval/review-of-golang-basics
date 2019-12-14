package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Hello")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Hello")

	}
}

func main() {

	go heavy()
	superHeavy()

	fmt.Println("Fin")
	// time.Sleep(time.Second * 5)
	select {}

}
