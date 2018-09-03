/*
Have the function SwapCase(str) take the str parameter 
and swap the case of each character. For example: if str 
is "Hello World" the output should be hELLO wORLD. 
Let numbers and symbols stay the way they are. 
*/

package main

import (
	"fmt"
	"unicode"
)

func SwapCase(str string) (result string) {

	//convert string to rune slice to use unicode package
	str2 := []rune(str) 

	for _, val := range str2 {
		if unicode.IsUpper(val){
			val = unicode.ToLower(val)
		} else if unicode.IsLower(val) {
			val = unicode.ToUpper(val)
		} else {
			val = val
		}
		result += string(val) 
	}
	return result
}

func main() {
	fmt.Println(SwapCase("Hello World!")) //hELLO wORLD!
}

