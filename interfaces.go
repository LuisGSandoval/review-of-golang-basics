package main

import "fmt"

// Vehicle is a general interface
// Enforcing style
type Vehicle interface {
	Drive()
	Stop()
}

// Chevrolet is a struct type of Car
type Chevrolet struct {
	Model string
}

// NewChevrolet is good
func NewChevrolet(arg string) Vehicle {
	return &Chevrolet{Model: arg}
}

// Drive function forces the car to start driving
func (c Chevrolet) Drive() {
	fmt.Println("...Let's hit the road ")
}

// Stop function forces the car to stop driving
func (c Chevrolet) Stop() {
	fmt.Println("...Stopping this car")
}

func main9() {

	car := NewChevrolet("2019")

	car.Drive()

}
