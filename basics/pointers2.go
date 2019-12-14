package main

import "fmt"

func changeValue(val *int) {
	*val = 1
}

func main10() {

	var number int = 10

	copyOne := number + 1

	copyTwo := number + 2

	fmt.Println(number)
	fmt.Println(copyOne)
	fmt.Println(copyTwo)

	// Cambiamos el valor inicial
	changeValue(&number)
	fmt.Println("Nuevamente probamos si cambian las vainas")

	fmt.Println(number)
	fmt.Println(copyOne)
	fmt.Println(copyTwo)

}
