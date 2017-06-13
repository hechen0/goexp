package exps

import "fmt"

func TestChannel(){
	achannel := make(chan int, 1)

	c(achannel)

	number := <-achannel
	fmt.Println(number)
}

func c(q chan<- int){
	go func(){ q <- 10 }()
}