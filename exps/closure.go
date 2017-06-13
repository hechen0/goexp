package exps

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Printf("%s, %v\n", p.name, &p)
}

func TestClosure() {

	data1 := []*field{{"one"}, {"two"}, {"three"}}
	data2 := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data1 {
		fmt.Println(&v)
		go v.print()
	}

	for _, v := range data2 {
		fmt.Println(&v)
		go v.print()
	}


	time.Sleep(3 * time.Second)
}
