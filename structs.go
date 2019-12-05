package main

// Car is a structs => data encapsulation
// A struct is an abstract datatype that groups together logicly related data
type Car struct {
	Brand string
	Model int
}

func main() {
	// Hay 2 formas de inicializar (agregar data en)  un struct la primer opcion es:
	c := Car{"Mazda", 2019} // Sin agregar el key de los valores
}
