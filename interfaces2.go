package main

import "fmt"

// An empty interface is written like this
// interface{} // This means that it can be of any type (string|int|boolean)

// An empty struct is written like this
// struct{}{} // this is useful when you reacieve a struct, but it's empty
// We can give values to a struct in an annonymous way too
// struct{Name,Age}{"Luis", 25} // This means that it can be of any type (string|int|boolean)

// Anything recieves anythin (string|int|boolean) and prints it out
func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main11() {
	Anything("Strings")
	Anything(true)
	Anything(5)
	Anything(5.6)
	Anything(struct{}{})

	// Creado maps
	myapp := make(map[string]interface{})

	myapp["name"] = "Luis"
	myapp["age"] = 25
	myapp["height"] = 1.73

	fmt.Println(myapp)

}
