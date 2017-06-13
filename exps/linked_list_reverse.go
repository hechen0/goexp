package exps

import (
	"fmt"
	"strconv"
)

type node struct {
	Value string
	Next *node
}

func ReverseLinkedList(){
	queue := NewLinkedList(5)
	printList(queue)

	current := queue
	var prev *node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	printList(prev)
}

func NewLinkedList(length int) *node {

	queue := new(node)
	queue.Value = "start"

	count := length
	current := queue
	for count > 0 {
		current.Next = &node{Value: strconv.Itoa(length - count + 1)}
		current = current.Next
		count--
	}
	return queue
}

func printList(list *node){
	for list != nil {
		fmt.Printf("%s-> ", list.Value)
		list  = list.Next
	}
	fmt.Println()
}