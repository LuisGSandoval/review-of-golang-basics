package main

import "fmt"

// Car is a structs => data encapsulation
// A struct is an abstract datatype that groups together logicly related data
type Car struct {
	Brand string
	Model int
}

// PrintCar es una funcion que se agrego al struct Car
func (c Car) PrintCar() {
	fmt.Println(c)
}

func main() {
	// Hay 2 formas de inicializar (agregar data en)  un struct la primer opcion es:
	// c := Car{"Mazda", 2019} // Sin agregar el key de los valores
	c := Car{
		Brand: "Mazda",
		Model: 2019,
	} // Agregando el key de los valores (recommended)

	// fmt.Println(c)
	c.PrintCar() // Estamos ejecutando la funcion que fue agregada a nuestro struct

}
