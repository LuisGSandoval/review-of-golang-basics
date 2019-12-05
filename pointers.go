package main

import (
	"fmt"
)

func swap(num1, num2 *int) {

	temp := *num2
	*num2 = *num1
	*num1 = temp

}

func main6() {

	// name := "Luis"
	// pointer := &name
	// fmt.Println("La direccion es: ", pointer)

	n1, n2 := 2, 30

	fmt.Println(n1, n2)
	swap(&n1, &n2)
	fmt.Println(n1, n2)

}
