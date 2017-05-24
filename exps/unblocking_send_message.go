package exps

import (
	"fmt"
)

func UnblockingSendMsg(){
	messages := make(chan string,1)
	msg := "hi"

	for i := 0; i < 2; i++ {
		select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		case hi := <- messages:
			fmt.Println("message get", hi)
		default:
		}
	}



}