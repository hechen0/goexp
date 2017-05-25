package exps

import (
	"runtime"
	"reflect"
	"fmt"
)



func GetFuncName(){
	var a func(string) error
	a = func(foo string) error {
		return nil
	}

	printname(a)
	printname(GetFuncName)
}

func printname(a interface{}){
	var fc *runtime.Func

	fc = runtime.FuncForPC(reflect.ValueOf(a).Pointer())
	fmt.Println(fc.Entry())
	fmt.Println(fc.FileLine(reflect.ValueOf(a).Pointer()))
}