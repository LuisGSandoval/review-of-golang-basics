package main

import "fmt"

func main() {

	// Continue, break
	flag := true
	for i := 0; i < 10; i++ {
		if i == 5 {
			flag = false
			break
		} else if i == 3 {
			continue
		}
		fmt.Println(i)

	}
	fmt.Println(flag)

	// Switch

	day := "Fri"
	switch day {
	case "Mon", "Tue", "Wed":
		fmt.Println("Working hard 😡")
	case "Thu":
		fmt.Println("Things are starting to get better son")
	case "Fri":
		fmt.Println("TGIF")
	default:
		fmt.Println("Oooopps an error has ocurred")
	}

}
