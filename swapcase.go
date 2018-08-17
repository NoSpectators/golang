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

	for _, val := range str {
		if unicode.IsUpper(rune(val)){
			val = unicode.ToLower(rune(val))
		} else if unicode.IsLower(rune(val)){
			val = unicode.ToUpper(rune(val))
		} else {
			val = val
		}
		result += string(val)		
	}
	return result
}

func main() {
	fmt.Println(SwapCase("Hello World!"))
}

