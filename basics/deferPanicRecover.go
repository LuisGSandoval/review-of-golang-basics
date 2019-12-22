package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("\n\n Defer, Panic & recover ")
	panicker()
	fmt.Println("Final de la execution")
}

func panicker() {
	fmt.Println("Listo para asustarse")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()

	panic("Algo malo sucedió")

}
