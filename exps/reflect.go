package exps

import (
	"runtime"
	"reflect"
	"fmt"
	"encoding/json"
)



func GetFuncName(){
	var a func(string) error
	a = func(foo string) error {
		return nil
	}

	printname(a)
	printname(GetFuncName)
}

func GetTypeName(){
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(result["status"]).Name())
}

func printname(a interface{}){
	var fc *runtime.Func

	fc = runtime.FuncForPC(reflect.ValueOf(a).Pointer())
	fmt.Println(fc.Entry())
	fmt.Println(fc.FileLine(reflect.ValueOf(a).Pointer()))
}