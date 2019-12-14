package main

import (
	"fmt"
	"os"
	"time"
)

func selectBlock(c chan int, exit chan int) {
	for {
		time.Sleep(time.Second)

		// Select works like an async switch but for channels only
		select {
		case <-c:
			fmt.Println("Recieved")
		case <-exit:
			fmt.Println("Exiting...")
			os.Exit(1)

		}
	}
}

func main() {

	c := make(chan int)
	exit := make(chan int)

	go func() {
		c <- 1
		close(c)
	}()

	go selectBlock(c, exit)

	time.Sleep(time.Second * 5)
	func() {
		exit <- 1
	}()

	select {}
}
