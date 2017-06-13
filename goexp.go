package main

import (
	"fmt"
)

func main() {
	TestChan()
}

func TestChan(){
	c := make(chan int)

	go func(){
		c <- 10
	}()

	fmt.Println(<-c)
}


