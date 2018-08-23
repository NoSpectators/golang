package main

import "fmt"

func Names() (first string, second string){
	first = "Foo"
	second = "Bar"
	return first, second
}

func main() {
	n1,n2 := Names()
	fmt.Println(n1, n2)
	n3, _ := Names()
	fmt.Println(n3) //skips second return value
}
