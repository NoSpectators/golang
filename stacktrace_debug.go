package main

import (
	"fmt"
	"reflect"
	"runtime" 
)

func sum13() {
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	fmt.Println("name of function is",GetFunctionName(sum13)) //6
}

