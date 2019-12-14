package main

import "fmt"

// in golang we have if else, for loops, switch, continue break, etc...
func main13() {
	// if else

	f := false
	flag := &f

	if flag == nil {
		fmt.Println("Nil value 😡")
	} else if *flag {
		fmt.Println("truthy value")
	} else {
		fmt.Println("Falsy value")
	}

	// For loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Infinite look
	// for {

	// }

	arr := []string{"Luis", "Maria", "Edilberto", "Antonio", "Pastor"}

	for _, v := range arr {
		fmt.Printf(" El nombre es: %v \n", v)
	}

	mymap := make(map[string]interface{})

	mymap["name"] = "Luis"
	mymap["age"] = 25

	for k, v := range mymap {
		fmt.Printf("%s = %v. \n", k, v)
	}

}
