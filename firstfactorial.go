package main


import (
	"fmt"
)



func FirstFactorial(num int) (result int) {

	//base case
	if num == 0{
		return 1
	} else { //recursive case
		return FirstFactorial(num - 1) * num
	}
	return result
}


// this is a comment
func main() {
    fmt.Println(FirstFactorial(5))
    fmt.Println(FirstFactorial(6))
}